#!/bin/sh

###########################################################
# precondition
###########################################################
#export ENC_KEY='xxxxxxxxxkeykey'
#export ENC_IV='xxxxxxxxxxxxiviv'

###########################################################
# Variable
###########################################################
#export GOTRACEBACK=single
GOTRACEBACK=all
CURRENTDIR=`pwd`

GO_LINT=1

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
# go lint
###########################################################
#go get -u github.com/golang/lint/golint
if [ $GO_LINT -eq 1 ]; then
    echo '============== golint =============='
    golint ./... | grep -v '^vendor\/' || true

    echo '============== misspell =============='
    #misspell .
    misspell `find . -name "*.go" | grep -v '/vendor/'`

    echo '============== ineffassign =============='
    ineffassign .
fi


###########################################################
# go list for check import package
###########################################################
#go list -f '{{.ImportPath}} -> {{join .Imports "\n"}}' ./xxxx.go


###########################################################
# go build and install
###########################################################
#build and install
#go build -i -v -o ./main ./gotestfile/main.go
go build -i -v -o ${GOPATH}/bin/gotestfile ./gotestfile/main.go

go build -i -v -o ${GOPATH}/bin/godepen ./godependency/main.go

go build -i -v -o ${GOPATH}/bin/gocipher ./gocipher/main.go

go build -i -v -o ${GOPATH}/bin/gobulkdata ./gobulkdata/main.go

###########################################################
# go run
###########################################################
#./main -n abc

###########################################################
# cross-compile for linux
###########################################################
#GOOS=linux go install -v ./...


###########################################################
# godoc
###########################################################
#godoc -http :8000
#http://localhost:8000/pkg/


###########################################################
# test
###########################################################
