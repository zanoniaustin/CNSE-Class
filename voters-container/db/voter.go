package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type VoterHistory struct {
	PollId   uint      `json:"poll_id"`
	VoteId   uint      `json:"vote_id"`
	VoteDate time.Time `json:"vote_date"`
}

type Voter struct {
	VoterId     uint           `json:"voter_id"`
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	VoteHistory []VoterHistory `json:"vote_history"`
}

type VoterList struct {
	cache
}

const (
	RedisNilError        = "redis: nil"
	RedisDefaultLocation = "0.0.0.0:6379"
	RedisKeyPrefix       = "todo:"
)

type cache struct {
	client  *redis.Client
	context context.Context
}

func New() (*VoterList, error) {
	redisUrl := os.Getenv("REDIS_URL")
	//This handles the default condition
	if redisUrl == "" {
		redisUrl = RedisDefaultLocation
	}
	return NewWithCacheInstance(redisUrl)
}

func NewWithCacheInstance(location string) (*VoterList, error) {

	//Connect to redis.  Other options can be provided, but the
	//defaults are OK
	client := redis.NewClient(&redis.Options{
		Addr: location,
	})

	//We use this context to coordinate betwen our go code and
	//the redis operaitons
	ctx := context.TODO()

	//This is the reccomended way to ensure that our redis connection
	//is working
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Println("Error connecting to redis" + err.Error())
		return nil, err
	}

	//Return a pointer to a new VoterList struct
	return &VoterList{
		cache: cache{
			client:  client,
			context: ctx,
		},
	}, nil
}

//------------------------------------------------------------
// REDIS HELPERS
//------------------------------------------------------------

func redisKeyFromVoterId(voterId uint) string {
	return fmt.Sprintf("%s%d", RedisKeyPrefix, voterId)
}

func (t *VoterList) getAllKeys() ([]string, error) {
	key := fmt.Sprintf("%s*", RedisKeyPrefix)
	return t.client.Keys(t.context, key).Result()
}

func fromJsonString(s string, voter *Voter) error {
	err := json.Unmarshal([]byte(s), &voter)
	if err != nil {
		return err
	}
	return nil
}

func (t *VoterList) upsertVoter(voter *Voter) error {
	log.Println("Adding new Id:", redisKeyFromVoterId(voter.VoterId))
	return t.client.JSONSet(t.context, redisKeyFromVoterId(voter.VoterId), ".", voter).Err()
}

func (t *VoterList) getVoterFromRedis(key string, voter *Voter) error {
	itemJson, err := t.client.JSONGet(t.context, key, ".").Result()
	if err != nil {
		return err
	}
	return fromJsonString(itemJson, voter)
}

func (t *VoterList) doesKeyExist(voterId uint) bool {
	kc, _ := t.client.Exists(t.context, redisKeyFromVoterId(voterId)).Result()
	return kc > 0
}

//------------------------------------------------------------
// PUBLIC FUNCTIONS TO SUPPORT VOTERS
//------------------------------------------------------------

func (vl *VoterList) AddVoter(voter *Voter) error {
	if vl.doesKeyExist(voter.VoterId) {
		return fmt.Errorf("Voter with voter id %d already exists", voter.VoterId)
	}
	return vl.upsertVoter(voter)
}

func (vl *VoterList) AddVoterHistory(id uint, voterHistory VoterHistory) error {
	voter, err := vl.GetVoter(id)
	if err != nil {
		return err
	}

	for _, poll := range voter.VoteHistory {
		if poll.PollId == voterHistory.PollId {
			return errors.New("poll id already exists")
		}
	}

	voterHistory.VoteDate = time.Now()
	voter.VoteHistory = append(voter.VoteHistory, voterHistory)

	return vl.upsertVoter(voter)
}

func (vl *VoterList) DeleteVoter(id uint) error {
	if !vl.doesKeyExist(id) {
		return fmt.Errorf("Voter with voter id %d does not exist", id)
	}

	return vl.client.Del(vl.context, redisKeyFromVoterId(id)).Err()
}

func (vl *VoterList) DeleteAll() (int, error) {
	keyList, err := vl.getAllKeys()
	if err != nil {
		return 0, err
	}

	numDeleted, err := vl.client.Del(vl.context, keyList...).Result()

	return int(numDeleted), err
}

func (vl *VoterList) UpdateVoter(updatedVoter *Voter) error {
	voter, err := vl.GetVoter(updatedVoter.VoterId)
	if err != nil {
		return fmt.Errorf("Voter with voter id %d does no exist", updatedVoter.VoterId)
	}
	updatedVoter.VoteHistory = voter.VoteHistory

	return vl.upsertVoter(updatedVoter)
}

func (vl *VoterList) GetVoter(id uint) (*Voter, error) {
	newVoter := &Voter{}
	err := vl.getVoterFromRedis(redisKeyFromVoterId(id), newVoter)
	if err != nil {
		return nil, err
	}

	return newVoter, nil
}

func (vl *VoterList) GetAllVoters() ([]Voter, error) {
	keyList, err := vl.getAllKeys()
	if err != nil {
		return nil, err
	}

	allVotersList := make([]Voter, len(keyList))
	for idx, key := range keyList {
		err := vl.getVoterFromRedis(key, &allVotersList[idx])
		if err != nil {
			return nil, err
		}
	}

	return allVotersList, nil
}

func (vl *VoterList) GetVoterHistory(id uint) ([]VoterHistory, error) {

	voter, err := vl.GetVoter(id)

	if err != nil {
		return nil, err
	}

	return voter.VoteHistory, nil
}

func (vl *VoterList) GetPollData(id uint, pollId uint) (VoterHistory, error) {

	voter, err := vl.GetVoter(id)
	if err != nil {
		return VoterHistory{}, err
	}

	for _, poll := range voter.VoteHistory {
		if poll.PollId == pollId {
			return poll, nil
		}
	}

	return VoterHistory{}, errors.New("poll data doesn't exist for voter")
}
