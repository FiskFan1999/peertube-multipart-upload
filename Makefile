# GIT_COMMIT and GIT_TAG borrowed from ergo open source project:
# https://raw.githubusercontent.com/ergochat/ergo/master/Makefile

GIT_COMMIT := $(shell git rev-parse HEAD 2> /dev/null)
GIT_TAG := $(shell git tag --points-at HEAD 2> /dev/null | head -n 1)

all: vers build

build:
	go build -v

install:
	go install -v

vers:
	echo '{"version":"$(GIT_TAG)", "commit": "$(GIT_COMMIT)"}' > version
