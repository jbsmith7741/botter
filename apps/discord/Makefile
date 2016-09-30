#!/usr/bin/env make

appName=discord
version=$(shell git describe --tags 2> /dev/null || git rev-parse HEAD)
internal="github.com/mediaFORGE/rta/internal"

remote: #loc=remote location
	GOOS=linux go build
	scp ./${appName} $(loc)
	rm ${appName}

