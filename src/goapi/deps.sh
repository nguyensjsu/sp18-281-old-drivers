#!/bin/bash

GOPATH=`pwd`
export GOPATH
go get github.com/go-redis/redis
go get github.com/satori/go.uuid
go get github.com/codegangsta/negroni
go get github.com/gorilla/mux
