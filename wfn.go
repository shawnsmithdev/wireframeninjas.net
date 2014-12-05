package main

import (
	"fmt"
	"net/http"
	"time"
)

func getTime() string {
	return time.Now().Truncate(time.Millisecond).Format(time.RFC3339Nano)
}

func getRobotsAllOK() string {
	return `User-agent: *
Disallow:
`
}

func getRoot() string {
	return `<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <style type="text/css">
      body {
        background-color: black;
        color: white;
      }
    </style>
    <script src="/wfn.js"></script>
    <title>WireframeNinjas</title>
  </head>
  <body>
  	<h4>Wireframe Ninjas</h4>
	<div>Coming soon: The return of a website nobody cares about!</div>
  	<canvas id="wfnCanvas" width="640" height="480"></canvas>
  </body>
</html>`
}

func getJS() string {
	return "console.log('Wireframe Ninjas!')";
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

func main() {
	http.HandleFunc("/", getHandler(getRoot))
	http.HandleFunc("/time", getHandler(getTime))
	http.HandleFunc("/wfn.js", getTypedHandler(getJS, "application/javascript"))
	http.HandleFunc("/robots.txt", getHandler(getRobotsAllOK))
	http.ListenAndServe(":8080", nil)
}
