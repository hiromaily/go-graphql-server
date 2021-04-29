# Note: tabs by space can't not used for Makefile!

CURRENTDIR=`pwd`
modVer=$(shell cat go.mod | head -n 3 | tail -n 1 | awk '{print $2}' | cut -d'.' -f2)
currentVer=$(shell go version | awk '{print $3}' | sed -e "s/go//" | cut -d'.' -f2)
gitTag=$(shell git tag | head -n 1)

###############################################################################
# Managing Dependencies
###############################################################################
.PHONY: check-ver
check-ver:
	#echo $(modVer)
	#echo $(currentVer)
	@if [ ${currentVer} -lt ${modVer} ]; then\
		echo go version ${modVer}++ is required but your go version is ${currentVer};\
	fi

.PHONY: update
update:
	GO111MODULE=off go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	go get -u -d -v ./...


###############################################################################
# Golang formatter and detection
###############################################################################
.PHONY: imports
imports:
	./scripts/imports.sh

.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: lintall
lintall: imports lint


###############################################################################
# Build
###############################################################################
.PHONY: build
build:
	go build -i -v -o ${GOPATH}/bin/graphql-server ./cmd/server/

.PHONY: build-version
build-version:
	go build -ldflags "-X main.version=${gitTag}" -i -v -o ${GOPATH}/bin/graphql-server ./cmd/server/

.PHONY: run
run:
	go run -v ./cmd/server/

