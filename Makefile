include .app
all: gen build test
build:
	go build -o $(APP_NAME) -v -ldflags " \
	-X magic-box/cmds.GitHash=$(shell git rev-parse HEAD 2> /dev/null) \
	-X magic-box/cmds.BuildTime=$(shell date "+%Y-%m-%dT%H:%M:%SZ")"
install:
	go install -ldflags " \
	-X magic-box/cmds.GitHash=$(shell git rev-parse HEAD 2> /dev/null) \
	-X magic-box/cmds.BuildTime=$(shell date "+%Y-%m-%dT%H:%M:%SZ")"
test:
	go test -v ./...
clean:
	go clean
	rm -f $(APP_NAME)
