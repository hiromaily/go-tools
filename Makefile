# Note: tabs by space can't not used for Makefile!

MONGO_PORT=27017
CURRENTDIR=`pwd`


###############################################################################
# Managing Dependencies
###############################################################################
.PHONY: update
update:
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u -d -v ./...


###############################################################################
# Golang formatter and detection
###############################################################################
.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: imports
imports:
	./scripts/imports.sh


###############################################################################
# cookie
#  show domain cookie
###############################################################################
.PHONY: run-cookie
run-cookie:
	go run -v ./cookie/main.go github.com


###############################################################################
# chat
###############################################################################
.PHONY: run-chat
run-chat:
	#localhost:8000
	go run -v ./chat/server/server.go
	go run -v ./chat/client/client.go


###############################################################################
# encryption
#  encrypt/decrypt
###############################################################################
.PHONY: run-encrypt
run-encrypt:
	go run -v ./encryption/main.go -m e secret-string

.PHONY: run-decrypt
run-decrypt:
	go run -v ./encryption/main.go -m d AMd/qKM1itq9ojh9nVEzDg==


###############################################################################
# gen-testfile
#  generate *_test.go template at current package(directory)
###############################################################################
.PHONY: run-gen-testfile
run-gen-testfile:
	go run -v ./gen-testfile/main.go -n new-package-name


###############################################################################
# gen-tls-cert
#  generate TLS cert file (cert.pem/key.pem)
###############################################################################
.PHONY: run-gen-tls-cert
run-gen-tls-cert:
	go run -v ./gen-tls-cert/main.go -host hy


###############################################################################
# gen-struct
#  generate go struct type from json data
###############################################################################
.PHONY: run-gen-struct
run-gen-struct:
	go run -v ./gen-struct/main.go -file ./json/teachers.json

.PHONY: run-gen-struct2
run-gen-struct2:
	go run -v ./gen-struct/main.go -json ./json/teachers.json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "date": "2017-07-26T11:10:15+02:00", "obj": {"child":100}}'


###############################################################################
# go-dependency
#  show dependencies of spacific package
#    cd ${GOPATH}/src/github.com/hiromaily/go-microservice
#    git checkout e1a9a740d0abfea5d190daa2d7b033799a06fe7f
#
#    cd ${GOPATH}/src/github.com/hiromaily/go-nats
#    git checkout 1f88d6d0063e2b3d6bba10208ad596905ac40dbb
#    ...
###############################################################################
.PHONY: run-go-dependency
run-go-dependency:
	go run -v ./go-dependency/main.go -target ${HOME}/work/go/src/github.com/hiromaily


###############################################################################
# substr
#  adjust time at srt file
###############################################################################
.PHONY: run-substr
run-substr:
	go run -v ./subsrt/main.go -f ./xxxxx.srt -t -1.5


###############################################################################
# Utility
###############################################################################
.PHONY: clean
clean:
	rm -rf cert.pem key.pem
