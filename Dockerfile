FROM golang:1.6

# Persistant folders (eventually)
RUN mkdir -p /srv/logs
RUN mkdir -p /srv/caddy

# Get dependencies
RUN go get github.com/mholt/caddy
RUN go get github.com/julienschmidt/httprouter

# Setup build
RUN mkdir -p /go/src/wfn
COPY wfn.go /go/src/wfn/
COPY wfn.sh /
COPY caddy /caddy
RUN ln -s /srv/caddy /caddy/.caddy

# Build backend
WORKDIR /go/src/wfn
RUN go install wfn

# Run container script
WORKDIR /caddy
ENTRYPOINT ["/wfn.sh"]

EXPOSE 80
# EXPOSE 443
