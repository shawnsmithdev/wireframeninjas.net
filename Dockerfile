FROM golang:1.6

RUN go get github.com/mholt/caddy
RUN go get github.com/julienschmidt/httprouter

COPY wfn.go /wfn.go
RUN go build -o /bin/wfn /wfn.go

COPY wfn.sh /wfn.sh
COPY caddy /caddy

RUN mkdir -p /home/caddy/logs
RUN mkdir -p /home/caddy/.caddy

ENV HOME /home/caddy
WORKDIR /home/caddy

ENTRYPOINT ["/wfn.sh"]

EXPOSE 80
