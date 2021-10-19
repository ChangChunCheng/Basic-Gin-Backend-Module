// Package main
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"basic-gin-backend-module/api"
	"basic-gin-backend-module/loader"
)

var (
	// BuildDate date string of when build was performed filled in by -X compile flag
	BuildDate string

	// BuiltOnIP date string of when build was performed filled in by -X compile flag
	BuiltOnIP string

	// BuiltOnOs date string of when build was performed filled in by -X compile flag
	BuiltOnOs string

	// RuntimeVer date string of when build was performed filled in by -X compile flag
	RuntimeVer string

	// LatestCommit date string of when build was performed filled in by -X compile flag
	LatestCommit string

	// BuildNumber date string of when build was performed filled in by -X compile flag
	BuildNumber string

	// OsSignal signal used to shutdown
	OsSignal chan os.Signal

	// AppPort
	// appPort string = mustGetenv("APP_PORT")
)

// GinServer - Build the Gin RestfulAPI Server
// return - gin.Router
// error - gin.Router init error
func GinServer() (err error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	api.BuildRouter(router)

	// router.Run(viper.GetString("APP_HOST") + ":" + viper.GetString("APP_PORT"))
	router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error starting server, the error is '%v'", err)
		panic("Gin init error.")
	}

	return
}

func main() {
	// Load System config and architecture
	loader.Load()

	OsSignal = make(chan os.Signal, 1)

	// Define version information
	fmt.Printf(
		`
		Application build information
		Build date      : %s
		Build on IP     : %s
		Built on OS     : %s
		Runtime version : %s
		Build number    : %s
		Git commit      : %s
		`,
		BuildDate, BuiltOnIP, RuntimeVer, BuiltOnOs, BuildNumber, LatestCommit,
	)

	go GinServer()
	LoopForever()
}

// LoopForever on signal processing
func LoopForever() {
	fmt.Println("Entering infinite loop")

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	_ = <-OsSignal

	fmt.Println("Exiting infinite loop received OsSignal")

}
