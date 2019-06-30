FROM golang:1.12.6-alpine3.10 AS build-env

RUN apk update && \
    apk add --no-cache git ca-certificates

ADD . /work/

WORKDIR /work

ENV GO111MODULE=on

RUN go mod download

CMD ["go", "test", "-cover", "-gcflags=-l", "./..."]