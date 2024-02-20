package api

import (
	"log"
	"net/http"
	"strconv"

	"drexel.edu/todo/db"
	"github.com/gin-gonic/gin"
)

type VotersAPI struct {
	db *db.VoterList
}

func New() (*VotersAPI, error) {
	dbHandler, err := db.New()
	if err != nil {
		return nil, err
	}

	return &VotersAPI{db: dbHandler}, nil
}

func (va *VotersAPI) ListAllVoters(c *gin.Context) {

	votersList, err := va.db.GetAllVoters()
	if err != nil {
		log.Println("Error Getting All Voters: ", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if votersList == nil {
		votersList = make([]db.Voter, 0)
	}

	c.JSON(http.StatusOK, votersList)
}

func (va *VotersAPI) GetVoter(c *gin.Context) {

	idString := c.Param("id")
	id64, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		log.Println("Error converting id to uint64: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	voter, err := va.db.GetVoter(uint(id64))
	if err != nil {
		log.Println("Voter not found: ", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, voter)
}

func (va *VotersAPI) GetVoterHistory(c *gin.Context) {
	idString := c.Param("id")
	id64, err := strconv.ParseUint(idString, 10, 32)
	if err != nil {
		log.Println("Error converting id to uint64: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	voterHistory, err := va.db.GetVoterHistory(uint(id64))
	if err != nil {
		log.Println("Voter History not found: ", err)
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, voterHistory)
}

// TODO
func (va *VotersAPI) GetPollData(c *gin.Context) {

	// idString := c.Param("id")
	// id64, err := strconv.ParseUint(idString, 10, 32)
	// if err != nil {
	// 	log.Println("Error converting id to uint64: ", err)
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }

	// voter, err := va.db.GetVoter(uint(id64))
	// if err != nil {
	// 	log.Println("Voter History not found: ", err)
	// 	c.AbortWithStatus(http.StatusNotFound)
	// 	return
	// }

	// pollIdString := c.Param("id")
	// pollId64, err := strconv.ParseUint(pollIdString, 10, 32)
	// if err != nil {
	// 	log.Println("Error converting pollId to uint64: ", err)
	// 	c.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }

	// pollData, err := va.db.GetPollData(voter, uint(pollId64))
	// if err != nil {
	// 	log.Println("Poll data not found: ", err)
	// 	c.AbortWithStatus(http.StatusNotFound)
	// 	return
	// }

	// c.JSON(http.StatusOK, pollData)
}

func (va *VotersAPI) AddVoter(c *gin.Context) {
	var voter db.Voter

	if err := c.ShouldBindJSON(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := va.db.AddVoter(voter); err != nil {
		log.Println("Error adding voter: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, voter)
}

// TODO
func (va *VotersAPI) AddPollData(c *gin.Context) {
	var voter db.Voter

	if err := c.ShouldBindJSON(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := va.db.AddVoter(voter); err != nil {
		log.Println("Error adding voter: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, voter)
}

func (va *VotersAPI) UpdateVoter(c *gin.Context) {
	var voter db.Voter
	if err := c.ShouldBindJSON(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := va.db.UpdateVoter(voter); err != nil {
		log.Println("Error updating voter: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, voter)
}

func (va *VotersAPI) DeleteVoter(c *gin.Context) {
	idString := c.Param("id")
	id64, _ := strconv.ParseUint(idString, 10, 32)

	if err := va.db.DeleteVoter(uint(id64)); err != nil {
		log.Println("Error deleting voter: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (va *VotersAPI) DeleteAllVoters(c *gin.Context) {

	if err := va.db.DeleteAll(); err != nil {
		log.Println("Error deleting all voters: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (va *VotersAPI) CrashSim(c *gin.Context) {
	//panic() is go's version of throwing an exception
	panic("Simulating an unexpected crash")
}

func (va *VotersAPI) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"status":             "ok",
			"version":            "1.0.0",
			"uptime":             100,
			"users_processed":    1000,
			"errors_encountered": 10,
		})
}
