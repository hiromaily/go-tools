# Note: tabs by space can't not used for Makefile!
MONGO_PORT=27017


###############################################################################
# Golang detection and formatter
###############################################################################
fmt:
	go fmt `go list ./... | grep -v '/vendor/'`

vet:
	go vet `go list ./... | grep -v '/vendor/'`

lint:
	golint ./... | grep -v '^vendor\/' || true
	misspell `find . -name "*.go" | grep -v '/vendor/'`
	ineffassign .

chk:
	go fmt `go list ./... | grep -v '/vendor/'`
	go vet `go list ./... | grep -v '/vendor/'`
	golint ./... | grep -v '^vendor\/' || true
	misspell `find . -name "*.go" | grep -v '/vendor/'`
	ineffassign .


###############################################################################
# Build
###############################################################################
bld1:
	go build -i -v -o ${GOPATH}/bin/gotestfile ./gotestfile/main.go

bld2:
	go build -i -v -o ${GOPATH}/bin/gocipher ./gocipher/main.go

bld3:
	go build -i -v -o ${GOPATH}/bin/godepen ./godependency/main.go

bld4:
	go build -i -v -o ${GOPATH}/bin/gobulkdata ./gobulkdata/main.go

bld5:
	go build -i -v -o ${GOPATH}/bin/servermain ./gochat/server/server.go
	go build -i -v -o ${GOPATH}/bin/clientmain ./gochat/client/client.go

bldall:bld1 bld2 bld3 bld4 bld5


###############################################################################
# Execution
###############################################################################
testpkg:
	gotestfile -n new-package-name

enc:
	gocipher -m e abcdefg

dec:
	gocipher -m d B4VmdhJuWkTXxyvTTDCG5w==

dep:
	godepen -target ${HOME}/work/go/src/github.com

bulk:
	gobulkdata -f ${HOME}/work/go/src/github.com/hiromaily/gotools/text.txt -l 20

chat:
	#localhost:8000
	./servermain &
	./clientmain

