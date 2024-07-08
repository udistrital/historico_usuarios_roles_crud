FROM golang:1.22.4-alpine
RUN apk add --no-cache git
RUN apk add --no-cache build-base
 
RUN go install github.com/beego/bee/v2@latest
WORKDIR /go/src