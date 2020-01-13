FROM golang:1.13-alpine as builder
ADD . /go/src/github.com/usedclothjp/grpc-nuxt-node-go-sample
WORKDIR /go/src/github.com/usedclothjp/grpc-nuxt-node-go-sample

ENV GO111MODULE=on

RUN apk update && \
    apk add git && \
    go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    GOOS=linux GOARCH=amd64 go build main.go

FROM alpine:3.9
WORKDIR /app
COPY --from=0 /go/src/github.com/usedclothjp/grpc-nuxt-node-go-sample /app

ENV PORT 80
EXPOSE 8080

CMD ["./main"]
