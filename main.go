package main

import (
	"fmt"
	"net/http"
	"time"
)

/* To build and run
	go build
	go run .
	open a browser on: http://127.0.0.1:9000
 */
func main() {
	// Register a few routes
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
		fmt.Fprintf(writer,"This is the help page. Time now: %s", time.Now().Format(time.RFC822))
	})

	// Handler to server a file system directory
	var fileSystem = http.Dir("C:\\Projects_Go\\SimpleHttpServer")
	var fileServer = http.FileServer(fileSystem);
	http.Handle("/files/", http.StripPrefix( "/files", fileServer ))		// requires a trailing /


	// Listen in on port 9000 on any address and use our simple handler
	// which implements the Handler interface
	http.ListenAndServe(":9000", nil)
}
