package api

import (
	"log"
	"net/http"

	"drexel.edu/todo/db"
	"github.com/gofiber/fiber/v2"
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

func (va *VotersAPI) ListAllVoters(c *fiber.Ctx) error {

	votersList, err := va.db.GetAllVoters()
	if err != nil {
		log.Println("Error Getting All Voters: ", err)
		return fiber.NewError(http.StatusNotFound,
			"Error Getting All Items")
	}

	if votersList == nil {
		votersList = make([]db.Voter, 0)
	}

	return c.JSON(votersList)
}

func (va *VotersAPI) GetVoter(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	voter, err := va.db.GetVoter(uint(id))
	if err != nil {
		log.Println("Voter not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	return c.JSON(voter)
}

func (va *VotersAPI) GetVoterHistory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	voterHistory, err := va.db.GetVoterHistory(uint(id))
	if err != nil {
		log.Println("Voter History not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	return c.JSON(voterHistory)
}

func (va *VotersAPI) GetPollData(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	pollId, err := c.ParamsInt("pollid")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	pollData, err := va.db.GetPollData(uint(id), uint(pollId))
	if err != nil {
		log.Println("Poll data not found: ", err)
		return fiber.NewError(http.StatusNotFound)
	}

	return c.JSON(pollData)
}

func (va *VotersAPI) AddVoter(c *fiber.Ctx) error {
	var voter db.Voter

	if err := c.BodyParser(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := va.db.AddVoter(&voter); err != nil {
		log.Println("Error adding voter: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(voter)
}

func (va *VotersAPI) AddPollData(c *fiber.Ctx) error {
	var voterHistory db.VoterHistory

	if err := c.BodyParser(&voterHistory); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := va.db.AddVoterHistory(uint(id), voterHistory); err != nil {
		log.Println("Error adding voter history: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(voterHistory)
}

func (va *VotersAPI) UpdateVoter(c *fiber.Ctx) error {
	var voter db.Voter

	if err := c.BodyParser(&voter); err != nil {
		log.Println("Error binding JSON: ", err)
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := va.db.UpdateVoter(&voter); err != nil {
		log.Println("Error updating voter: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.JSON(voter)
}

func (va *VotersAPI) DeleteVoter(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(http.StatusBadRequest)
	}

	if err := va.db.DeleteVoter(uint(id)); err != nil {
		log.Println("Error deleting voter: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).SendString("Delete OK")
}

func (va *VotersAPI) DeleteAllVoters(c *fiber.Ctx) error {

	if cnt, err := va.db.DeleteAll(); err != nil {
		log.Println("Error deleting all voters: ", err)
		return fiber.NewError(http.StatusInternalServerError)
	} else {
		log.Println("Deleted ", cnt, " items")
	}

	return c.Status(http.StatusOK).SendString("Delete All OK")
}

func (va *VotersAPI) CrashSim(c *fiber.Ctx) error {
	//panic() is go's version of throwing an exception
	panic("Simulating an unexpected crash")
}

func (va *VotersAPI) HealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"status":             "ok",
			"version":            "1.0.0",
			"uptime":             100,
			"users_processed":    1000,
			"errors_encountered": 10,
		})
}
