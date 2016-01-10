#!/bin/sh

# This is a script I use to actually build and push the prod docker image for wfn
# to AWS Elastic Container Registry.  You almost certainly don't care about this.

echo "Did you remember to update your docker login creds?";
echo "aws --region=us-east-1 ecr get-login";

go get -u github.com/julienschmidt/httprouter
go get -u github.com/mholt/caddy;

export CGO_ENABLED=0;
export GOOS=linux;

go build -ldflags "-s" -a -installsuffix cgo -o wfn wfn.go \
&& go build -ldflags "-s" -a -installsuffix cgo -o caddy-bin github.com/mholt/caddy \
&& sudo docker build -f ./Dockerfile.prod -t shawnsmithdev/wfn . \
&& sudo docker tag -f shawnsmithdev/wfn:latest 671121103357.dkr.ecr.us-east-1.amazonaws.com/shawnsmithdev/wfn:latest \
&& sudo docker push 671121103357.dkr.ecr.us-east-1.amazonaws.com/shawnsmithdev/wfn:latest
