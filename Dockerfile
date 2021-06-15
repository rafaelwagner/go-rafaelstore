FROM golang:latest

ADD . /go/src/gitlab.com/store/

WORKDIR /go/src/gitlab.com/store/

RUN go mod download

EXPOSE 8000