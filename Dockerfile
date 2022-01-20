FROM golang:lastest

RUN mkdir/build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get