FROM golang:latest

RUN mkdir -p /go/src/github.com/riceo/demo-app
WORKDIR /go/src/github.com/riceo/demo-app
COPY main.go ./
COPY vendor ./vendor/
RUN go get -u github.com/kardianos/govendor
RUN go build main.go
