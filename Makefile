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
run:bld6 srt


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

bld6:
	go build -i -v -o ${GOPATH}/bin/gosubsrt ./gosubsrt/main.go

bld7:
	go build -i -v -o ${GOPATH}/bin/gocookie ./gocookie/main.go

bld8:
	go build -i -v -o ${GOPATH}/bin/goapitest ./goapitest/main.go

bld9:
	go build -i -v -o ${GOPATH}/bin/gogentype ./gogentype/main.go

bldall:bld1 bld2 bld3 bld4 bld5 bld6 bld7 bld8


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

srt:
	gosubsrt -f ${HOME}/work/go/src/github.com/hiromaily/gotools/gosubsrt/srtfiles/Silicon.Valley.S02E01.srt -t 6.2

cookie:
	gocookie localhost
	#gocookie gist.github.com

apitest:
	goapitest -m user

gentype:
	gogentype -json '{"str": "xxxx", "slice": [1,2,3], "sliceempty": [], "null": null, "int": 10, "zero": 0, "bool": true, "date": "2017-07-26T11:10:15+02:00", "obj": {"child":100}}'

gentype2:
	gogentype -file $(PWD)/gogentype/json/teachers.json
