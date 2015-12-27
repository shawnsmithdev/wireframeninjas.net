wireframeninjas.net
===================

Application that runs on wireframeninjas.net  Major components are:

* wfn.go - A go web app for dynamic content. Serves content to the Caddy proxy.
* caddy/Caddyfile - Configuration for Caddy, a HTTP/2 complient frontend
  for TLS termination and static content.  It proxies dynamic content to wfn.go.
* caddy/root - Static HTML and such for frontend.
* Dockerfile - Builds the Docker container for building and running the website.
* wfn.sh - A script that runs the front and backends within the docker container.

You can run it like this:
    docker build -t wfn-test .
    docker run -it -p 80:80 wfn-test
