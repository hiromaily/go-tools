#!/bin/sh

###########################################################
# Variable
###########################################################
#export GOTRACEBACK=single
GOTRACEBACK=all
CURRENTDIR=`pwd`

###########################################################
# Update all package
###########################################################
#go get -u -v ./...


###########################################################
# Adjust version dependency of projects
###########################################################
#cd ${GOPATH}/src/github.com/aws/aws-sdk-go
#git checkout v0.9.17
#git checkout master


###########################################################
# go fmt and go vet
###########################################################
go fmt ./...
go vet ./...
#go vet `go list ./... | grep -v '/vendor/'`


###########################################################
# go build and install
###########################################################
#build and install
go build -i -v -o ./servermain ./server/server.go
go build -i -v -o ./clientmain ./client/client.go


#localhost:8000
#./servermain &

#./clientmain

###########################################################
# test
###########################################################
#stress test
#https://github.com/rakyll/boom
# $ boom -n 1000 -c 100 https://google.com