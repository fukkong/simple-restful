#!/usr/bin/env bash

set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

go get github.com/gorilla/mux
go get github.com/gorilla/handlers

go build -o bin/application application.go