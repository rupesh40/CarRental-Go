FROM golang:1.18.0

WORKDIR /usr/src/app

copy . . 

RUN go mod tidy


