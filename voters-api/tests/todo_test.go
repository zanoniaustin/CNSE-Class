package tests

import (
	"log"
	"os"
	"testing"

	"drexel.edu/todo/db"
	fake "github.com/brianvoe/gofakeit/v6" //aliasing package name
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	BASE_API = "http://localhost:1080"

	cli = resty.New()
)

func TestMain(m *testing.M) {

	//SETUP GOES FIRST
	rsp, err := cli.R().Delete(BASE_API + "/voters")

	if rsp.StatusCode() != 200 {
		log.Printf("error clearing database, %v", err)
		os.Exit(1)
	}

	code := m.Run()

	//CLEANUP
	rsp1, err1 := cli.R().Delete(BASE_API + "/voters")
	if rsp1.StatusCode() != 200 {
		log.Printf("error cleaning up database, %v", err1)
		os.Exit(1)
	}

	//Now Exit
	os.Exit(code)
}

func newRandVoter(id uint) db.Voter {
	return db.Voter{
		VoterId: id,
		Name:    fake.Name(),
		Email:   fake.Email(),
	}
}

func newRandVoterHistory(id uint) db.VoterHistory {
	return db.VoterHistory{
		PollId: id,
		VoteId: fake.UintRange(0, 100),
	}
}

func Test_LoadDB(t *testing.T) {
	numLoad := 3
	for i := 0; i < numLoad; i++ {
		item := newRandVoter(uint(i))
		rsp, err := cli.R().
			SetBody(item).
			Post(BASE_API + "/voters")

		assert.Nil(t, err)
		assert.Equal(t, 200, rsp.StatusCode())
	}
}

func Test_GetAllVoters(t *testing.T) {
	var items []db.Voter

	rsp, err := cli.R().SetResult(&items).Get(BASE_API + "/voters")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	assert.Equal(t, 3, len(items))
}

func Test_GetVoter(t *testing.T) {
	var item db.Voter

	rsp, err := cli.R().SetResult(&item).Get(BASE_API + "/voters/1")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	assert.Equal(t, uint(1), item.VoterId)
}

func Test_LoadVoterHistory(t *testing.T) {
	numLoad := 3
	for i := 0; i < numLoad; i++ {
		item := newRandVoterHistory(uint(i))
		rsp, err := cli.R().
			SetBody(item).
			Post(BASE_API + "/voters/1/polls")

		assert.Nil(t, err)
		assert.Equal(t, 200, rsp.StatusCode())
	}
}

func Test_GetAllVoterHistory(t *testing.T) {
	var items []db.VoterHistory

	rsp, err := cli.R().SetResult(&items).Get(BASE_API + "/voters/1/polls")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	assert.Equal(t, 3, len(items))
}

func Test_GetVoterHistory(t *testing.T) {
	var item db.VoterHistory

	rsp, err := cli.R().SetResult(&item).Get(BASE_API + "/voters/1/polls/1")

	assert.Nil(t, err)
	assert.Equal(t, 200, rsp.StatusCode())

	assert.Equal(t, uint(1), item.PollId)
}
