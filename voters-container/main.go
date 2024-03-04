package main

import (
	"flag"
	"fmt"
	"os"

	"drexel.edu/todo/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Global variables to hold the command line flags to drive the todo CLI
// application
var (
	hostFlag string
	portFlag uint
)

// processCmdLineFlags parses the command line flags for our CLI
//
// TODO: This function uses the flag package to parse the command line
//		 flags.  The flag package is not very flexible and can lead to
//		 some confusing code.

//			 REQUIRED:     Study the code below, and make sure you understand
//						   how it works.  Go online and readup on how the
//						   flag package works.  Then, write a nice comment
//				  		   block to document this function that highights that
//						   you understand how it works.
//
//			 EXTRA CREDIT: The best CLI and command line processor for
//						   go is called Cobra.  Refactor this function to
//						   use it.  See github.com/spf13/cobra for information
//						   on how to use it.
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
	r := gin.Default()
	r.Use(cors.Default())

	apiHandler, err := api.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r.GET("/voters", apiHandler.ListAllVoters)
	r.POST("/voters", apiHandler.AddVoter)
	r.PUT("/voters", apiHandler.UpdateVoter)
	r.DELETE("/voters", apiHandler.DeleteAllVoters)
	r.DELETE("/voters/:id", apiHandler.DeleteVoter)
	r.GET("/voters/:id", apiHandler.GetVoter)
	r.GET("/voters/:id/polls", apiHandler.GetVoterHistory)
	r.POST("/voters/:id/polls", apiHandler.AddPollData)
	r.GET("/voters/:id/polls/:pollid", apiHandler.GetPollData)

	r.GET("/crash", apiHandler.CrashSim)
	r.GET("/health", apiHandler.HealthCheck)

	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	r.Run(serverPath)
}
