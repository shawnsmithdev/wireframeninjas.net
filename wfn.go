package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

const (
	// Address to serve the backend on
	serveAddress string = ":8081"
	// Sleep for just /slow in ms
	sleepDefaultMillis int = 250
	// Format for reporting on sleep
	sleepJsonFmt string = `{"sleep_unit": "Millisecond", "sleep_amount", %d}`
	// If anyone could put 2^31-1 in, someone would, eventually.
	maxSleepMs int = 5000
)

// Handle /time for timestamps
func timeHandler(resp http.ResponseWriter, req *http.Request) {
	withType("text/plain", handleString(func() string {
		return time.Now().Truncate(time.Millisecond).Format(time.RFC3339Nano)
	}))(resp, req)
}

// Builds a handler that sleeps for the given number of ms and returns a json about it.
func handleSleep(ms int) http.HandlerFunc {
	sleepTime := time.Duration(ms) * time.Millisecond
	return withType("application/json", handleString(func() string {
		time.Sleep(sleepTime)
		return fmt.Sprintf(sleepJsonFmt, ms)
	}))
}

// Handle /slow and /slow/[int], simulate slow rest calls.
func slowHandler(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	sleepStr := params.ByName("sleep")
	if sleepStr == "" {
		handleSleep(sleepDefaultMillis)(resp, req)
		return
	}
	sleep, err := strconv.Atoi(sleepStr)
	if err != nil {
		handleString(err.Error)(resp, req)
		return
	}
	if sleep > maxSleepMs {
		sleep = maxSleepMs
	}
	handleSleep(sleep)(resp, req)
}

// Build simple string handlers
func handleString(stringSource func() string) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(resp, stringSource())
	}
}

// Inject content type headers
func withType(contentType string, wrapped http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", contentType)
		wrapped(resp, req)
	}
}

// Adaptor to ignore parameters
func noParams(wrapped http.HandlerFunc) httprouter.Handle {
	return func(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		wrapped(resp, req)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/time", noParams(timeHandler))
	router.GET("/slow", slowHandler)
	router.GET("/slow/:sleep", slowHandler)

	fmt.Println(http.ListenAndServe(serveAddress, router))
}
