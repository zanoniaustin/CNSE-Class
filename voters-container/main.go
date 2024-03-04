package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"drexel.edu/todo/api"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Global variables to hold the command line flags to drive the todo CLI
// application
var (
	hostFlag string
	portFlag uint
)

// processCmdLineFlags parses the command line flags for our CLI
//
//	 YOUR ANSWER: Command line flag options:
//					-h set the IP address to listen on, default 0.0.0.0
//					-p set the port that service will connect to, default 1080
//				  Flags are first initially declared, then flag.Parse() is ran to
//				  see which flags are present and if they have any arguments
func processCmdLineFlags() {

	flag.StringVar(&hostFlag, "h", "0.0.0.0", "Listen on all interfaces")
	flag.UintVar(&portFlag, "p", 1080, "Default Port")

	flag.Parse()
}

func main() {
	processCmdLineFlags()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())

	apiHandler, err := api.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app.Get("/voters", apiHandler.ListAllVoters)
	app.Post("/voters", apiHandler.AddVoter)
	app.Put("/voters", apiHandler.UpdateVoter)
	app.Delete("/voters", apiHandler.DeleteAllVoters)
	app.Delete("/voters/:id", apiHandler.DeleteVoter)
	app.Get("/voters/:id", apiHandler.GetVoter)
	app.Get("/voters/:id/polls", apiHandler.GetVoterHistory)
	app.Post("/voters/:id/polls", apiHandler.AddPollData)
	app.Get("/voters/:id/polls/:pollid", apiHandler.GetPollData)

	app.Get("/crash", apiHandler.CrashSim)
	app.Get("/health", apiHandler.HealthCheck)

	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	log.Println("Starting server on ", serverPath)
	app.Listen(serverPath)
}
