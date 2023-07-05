FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

