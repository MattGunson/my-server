package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MattGunson/my-server/api"
)

const (
	defaultPort    = "3001"
	defaultLogFile = "/var/log/svs_access.json"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// logFilePath := os.Getenv("LOG_FILE_PATH")
	// if logFilePath == "" {
	// 	logFilePath = defaultLogFile
	// }

	// logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	// if err != nil {
	// 	panic(err)
	// }
	// defer logFile.Close()

	api, err := api.NewAPI()
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":"+defaultPort, api))
}
