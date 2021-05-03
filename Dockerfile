#!/bin/bash

FROM golang:1.16.3

ENV GO111MODULE=on


WORKDIR /cmd

COPY . .

RUN apt-get update
RUN apt-get -y install postgresql-client


RUN go build -o app-constructor-backend ./cmd/main.go

CMD ["./app-constructor-backend"]

