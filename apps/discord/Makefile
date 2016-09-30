#!/usr/bin/env make

appName=discord
version=$(shell git describe --tags 2> /dev/null || git rev-parse HEAD)
internal="github.com/jbsmith7741/botter/internal"

build:
	go build -ldflags "-X ${internal}/version.AppName=${appName} -X ${internal}/version.Version=${version} -X ${internal}/version.BuildTimeUTC=`date -u '+%Y-%m-%d_%I:%M:%S%p'`"

remote: #loc=remote location
	GOOS=linux make build
	scp ./${appName} $(loc)
	rm ${appName}

