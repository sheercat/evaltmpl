#!/bin/sh

AP=evaltmpl

GOOS=linux GOARCH=amd64 go build $AP.go
docker build -f Dockerfile.local -t $AP .
docker run -it --rm -p 3000:3000 $AP

