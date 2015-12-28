wireframeninjas.net
===================

Application that runs on [wireframeninjas.net](wireframeninjas.net). Major components are:

* caddy/Caddyfile - Configuration for [Caddy](https://caddyserver.com/), an HTTP/2
  capable frontend for TLS termination and static content.  It proxies dynamic content
  to the backend over port 8081 inside the container.
* caddy/root - Static HTML and such for the frontend.
* wfn.go - A go web app for the dynamic content backend.
* Dockerfile - Builds a Docker container for building and running the website.
* Dockerfile.min - Builds a Docker container just for running the website.
* Dockerfile.prod.min - (eventually) Builds a Docker container with automatic TLS through Let's Encrypt.
* wfn.sh - A script that runs the frontend and backend within the docker container.

You can run it like this:
```
    docker build -t wfn-test .
    docker run -it -p 80:80 wfn-test
```

You can also build a smaller docker image with `Dockerfile.min`
```
  go get github.com/julienschmidt/httprouter
  go get github.com/mholt/caddy
  CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o caddy-bin github.com/mholt/caddy
  CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo -o wfn wfn.go
  docker build -f ./Dockerfile.min -t wfn-min . && docker run -it -p 80:80 wfn-min
```

There will also eventually be a `Dockerfile.prod.min` that will have automatic TLS through Let's Encrypt.
You won't want to use it unless have access to the wireframeninjas.net domain because you are me.
