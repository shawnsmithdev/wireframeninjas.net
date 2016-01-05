wireframeninjas.net
===================

Application that runs on [wireframeninjas.net](http://wireframeninjas.net). Major components are:

* caddy/Caddyfile - Configuration for [Caddy](https://caddyserver.com/), an HTTP/2 capable frontend for TLS
  termination and static content.  It proxies dynamic content to the backend over port 8081 inside the container.
* caddy/root - Static HTML and such for the frontend.
* wfn.go - A go web app for the dynamic content backend.
* Dockerfile - Builds a Docker container for building and running the website.
* Dockerfile.prod - Builds a minimal alpine-based Docker container with automatic TLS through Let's Encrypt.
  You won't want to use this unless have access to the wireframeninjas.net domain because *you are me*.
* wfn.sh - A script that runs the frontend and backend within the docker container.
* wfn.prod.sh - The same script, except it calls Caddy with arguments to support automatic TLS. Requires an
  environmental variable to feed it the registration email.

You can run it like this:
```
    docker build -t wfn-test .
    docker run -it -p 80:80 wfn-test
```


#### Bugs and Gotchas

There's [a problem](https://github.com/letsencrypt/boulder/issues/1279) with Let's Encrypt's API when used
over HTTP/2, so for now you should build Caddy with Go 1.5.

I'm running this on [AWS ECS](https://aws.amazon.com/ecs/), but I'm using volumes for the caddy files that
include the private certificate. The volumes are stored in
[an encrypted EBS volume](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) out of excessive
caution.  However, you need to ensure you reboot after adding the EBS mount to `/etc/fstab` so that the Docker
daemon will see it, otherwise the files will never escape the docker container.

Alpine doesn't have CA certs installed by default, so we needed to add it so we connect to Let's Encrypt over TLS.
