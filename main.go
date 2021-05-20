package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

/* To build and run
go build
go run .
open a browser on: http://127.0.0.1:9000
*/
func main() {
	// Create a wait group (count down) to wait for two http servers to terminate
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	// Register a few routes (shared by both server, for now)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// Get the response header
		header := writer.Header()

		// Let browser know we are sending back JSON
		header.Set("Content-Type", "application/json")

		// Reset header data
		header.Set("Date", time.Now().Format(time.RFC3339))

		// Reset header status
		writer.WriteHeader(http.StatusOK)

		// Respond with a JSON string
		fmt.Fprint(writer, `{"status":"OK"}`)
	})
	http.HandleFunc("/help", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "This is the help page. Time now: %s", time.Now().Format(time.RFC822))
	})

	// Handler to server a file system directory
	var fileSystem = http.Dir("C:\\Projects_Go\\SimpleHttpServer")
	var fileServer = http.FileServer(fileSystem)
	http.Handle("/files/", http.StripPrefix("/files", fileServer)) // requires a trailing /

	// Start one server
	go func() {
		// Listen in on port 9000 on any address
		err := http.ListenAndServe(":9000", nil)
		if err != nil {
			log.Printf("Server at port 9000 failed: %s\n", err.Error())
		}
		waitGroup.Done()
	}()

	// Start another server
	go func() {
		// Listen in on port 9000 on any address
		err := http.ListenAndServe(":9001", nil)
		if err != nil {
			log.Printf("Server at port 9001 failed: %s\n", err.Error())
		}
		waitGroup.Done()
	}()

	// Wait for both servers to terminate
	waitGroup.Wait()
}
