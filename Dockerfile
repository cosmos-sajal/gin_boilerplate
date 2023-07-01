FROM golang:alpine

RUN apk update && apk add --no-cache git

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

EXPOSE 3000

CMD CompileDaemon --build="go build main.go" --command=./main
