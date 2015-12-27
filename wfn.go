package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

const (
	sleepJsonFmt string = `{"sleep_unit": "Millisecond", "sleep_amount", %d}`
	maxSleepMs   int    = 5000
)

func getTime() string {
	return time.Now().Truncate(time.Millisecond).Format(time.RFC3339Nano)
}

func getHandler(stringGetter func() string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%v] Got Request from [%v]: [%v] %v%v \n", getTime(), r.RemoteAddr, r.Method, r.Host, r.URL)
		fmt.Fprintf(w, stringGetter())
	}
}

func getTypedHandler(stringGetter func() string, contentType string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		getHandler(stringGetter)(w, r)
	}
}

func getSlow(ms int) func() string {
	sleepTime := time.Duration(ms) * time.Millisecond
	return func() string {
		time.Sleep(sleepTime)
		return fmt.Sprintf(sleepJsonFmt, ms)
	}
}

func slowHandler(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	sleepStr := params.ByName("sleep")
	sleep, err := strconv.Atoi(sleepStr)
	if err != nil {
		getHandler(err.Error)(resp, req)
	} else {
		if sleep > maxSleepMs {
			sleep = maxSleepMs
		}
		getTypedHandler(getSlow(sleep), "application/json")(resp, req)
	}
}

func noParams(wrapped func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		wrapped(w, r)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/time", noParams(getHandler(getTime)))
	router.GET("/slow", noParams(getTypedHandler(getSlow(250), "application/json")))
	router.GET("/slow/:sleep", slowHandler)

	fmt.Println(http.ListenAndServe(":8081", router))
}
