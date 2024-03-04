package db

import (
	"errors"
	"time"
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
	Voters map[uint]Voter `json:"voter"`
}

func New() (*VoterList, error) {
	voterList := &VoterList{
		Voters: make(map[uint]Voter),
	}

	return voterList, nil
}

func (vl *VoterList) AddVoter(voter Voter) error {

	_, ok := vl.Voters[voter.VoterId]
	if ok {
		return errors.New("voter already exists")
	}

	voter.VoteHistory = []VoterHistory{}

	vl.Voters[voter.VoterId] = voter

	return nil
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
	vl.Voters[voter.VoterId] = voter

	return nil
}

func (vl *VoterList) DeleteVoter(id uint) error {

	_, ok := vl.Voters[id]
	if !ok {
		return errors.New("voter doesn't exist")
	}

	delete(vl.Voters, id)

	return nil
}

func (vl *VoterList) DeleteAll() error {
	vl.Voters = make(map[uint]Voter)

	return nil
}

func (vl *VoterList) UpdateVoter(voter Voter) error {

	_, ok := vl.Voters[voter.VoterId]
	if !ok {
		return errors.New("voter does not exist")
	}

	voter.VoteHistory = vl.Voters[voter.VoterId].VoteHistory

	vl.Voters[voter.VoterId] = voter

	return nil
}

func (vl *VoterList) GetVoter(id uint) (Voter, error) {

	voter, ok := vl.Voters[id]
	if !ok {
		return Voter{}, errors.New("voter does not exist")
	}

	return voter, nil
}

func (vl *VoterList) GetAllVoters() ([]Voter, error) {
	var voterList []Voter
	for _, item := range vl.Voters {
		voterList = append(voterList, item)
	}

	return voterList, nil
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
