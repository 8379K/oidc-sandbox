FROM golang:1.17-alpine

RUN apk update && apk add git

RUN mkdir /go/src/app

WORKDIR /8379K/oidc-sandbox

COPY . /8379K/oidc-sandbox/

RUN go mod download
