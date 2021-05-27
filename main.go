package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func createServer(port int) *http.Server {
	// Create a new ServeMux: Recall that a ServeMux is an HTTP request multiplexer.
	// It matches the URL of each incoming request against a list of registered patterns
	//and calls the handler for the pattern that most closely matches the URL.
	mux := http.NewServeMux()

	// Register a handle for the root
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
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

	mux.HandleFunc("/help", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "This is the help page. Time now: %s", time.Now().Format(time.RFC822))
	})

	mux.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		// Request some data from a remove server
		res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

		// Check for errors
		if err != nil {
			fmt.Fprintf(writer, "Error: %v", err)
		} else {
			// Read the response body
			data, _ := ioutil.ReadAll(res.Body)

			// Get the value of the Content-Type header
			contentType := res.Header.Get("Content-Type")

			// Close the response body and send data back
			res.Body.Close()
			fmt.Fprintf(writer, "%s\n", data)
			fmt.Fprintf(writer, "%s\n", contentType)

		}
	})

	// Create a new http server by initializing a Server struct with appropriate parameters
	//for running an HTTP server. The zero value for Server is a valid configuration
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	return &server
}

/* To build and run
go build
go run .
open a browser on: http://127.0.0.1:9000
*/
func main() {
	// Create a wait group (count down) to wait for two http servers to terminate
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	// Start one server
	go func() {
		server := createServer(9000)
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server at port 9000 failed: %s\n", err.Error())
		}
		waitGroup.Done()
	}()

	// Start another server
	go func() {
		server := createServer(9001)
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Server at port 9001 failed: %s\n", err.Error())
		}
		waitGroup.Done()
	}()

	// Wait for both servers to terminate
	waitGroup.Wait()
	fmt.Println("Both server terminated")
}
