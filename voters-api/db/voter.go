package db

import (
	"encoding/json"
	"errors"
	"fmt"
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

	vl.Voters[voter.VoterId] = voter

	//TODO: Fix Voter History maybe?

	return nil
}

func (vl *VoterList) GetVoter(id uint) (Voter, error) {

	voter, ok := vl.Voters[id]
	if !ok {
		return Voter{}, errors.New("voter does not exist")
	}

	return voter, nil
}

func (vl *VoterList) GetVoterHistory(id uint) ([]VoterHistory, error) {

	voter, err := vl.GetVoter(id)

	if err != nil {
		return nil, err
	}

	return voter.VoteHistory, nil
}

// ChangeItemDoneStatus accepts an item id and a boolean status.
// It returns an error if the status could not be updated for any
// reason.  For example, the item itself does not exist, or an
// IO error trying to save the updated status.

// Preconditions:   (1) The database file must exist and be a valid
//
//					(2) The item must exist in the DB
//	    				because we use the item.Id as the key, this
//						function must check if the item already
//	    				exists in the DB, if not, return an error
//
// Postconditions:
//
//	 (1) The items status in the database will be updated
//		(2) If there is an error, it will be returned.
//		(3) This function MUST use existing functionality for most of its
//			work.  For example, it should call GetItem() to get the item
//			from the DB, then it should call UpdateItem() to update the
//			item in the DB (after the status is changed).
func (vl *VoterList) ChangeItemDoneStatus(id int, value bool) error {

	//update was successful
	return errors.New("not implemented")
}

func (vl *VoterList) GetAllVoters() ([]Voter, error) {
	var voterList []Voter
	for _, item := range vl.Voters {
		voterList = append(voterList, item)
	}

	return voterList, nil
}

// PrintItem accepts a ToDoItem and prints it to the console
// in a JSON pretty format. As some help, look at the
// json.MarshalIndent() function from our in class go tutorial.
func (vl *VoterList) PrintItem(item VoterHistory) {
	jsonBytes, _ := json.MarshalIndent(item, "", "  ")
	fmt.Println(string(jsonBytes))
}

// PrintAllItems accepts a slice of ToDoItems and prints them to the console
// in a JSON pretty format.  It should call PrintItem() to print each item
// versus repeating the code.
func (vl *VoterList) PrintAllItems(itemList []VoterHistory) {
	for _, item := range itemList {
		vl.PrintItem(item)
	}
}

// JsonToItem accepts a json string and returns a ToDoItem
// This is helpful because the CLI accepts todo items for insertion
// and updates in JSON format.  We need to convert it to a ToDoItem
// struct to perform any operations on it.
func (vl *VoterList) JsonToItem(jsonString string) (VoterHistory, error) {
	var item VoterHistory
	err := json.Unmarshal([]byte(jsonString), &item)
	if err != nil {
		return VoterHistory{}, err
	}

	return item, nil
}
