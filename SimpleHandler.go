package main

import (
	"fmt"
	"net/http"
	"time"
)

// SimpleHandler implements the Handler interface
type SimpleHandler struct {}

// This method gets invoked for all incoming requests. The response gets
// sent back to the client when this method returns
func (sh SimpleHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	// Create response binary data
	bytes := []byte("Hello World")

	// Write the binary data to the response object using ResponseWriter.write method:
	// 	Write([]byte) (int, error) --> returns the number of bytes written to
	// the response and an error if Write() call fails.
	// We can simply ignore these values but the number of bytes written can be
	// useful to send Content-Length response header.
	res.Write(bytes)

	// Since both http.ResponseWriter and io.Writer interface implement the same
	// Write([]byte) (int, error), we can treat res as an io.Writer and use
	// any method that gets passed an io.Writer
	fmt.Fprintf(res, "Time now: %s", time.Now().Format("RFC822"))
}


