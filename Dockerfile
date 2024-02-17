# FROM golang:1.21.6


# COPY ./ .

# WORKDIR /www/go-jwt/cmd

# RUN go build -o ./bin/go-jwt

# CMD ["./bin/go-jwt"]

FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/cespare/reflex

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ././bin/go-jwt

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./bin/go-jwt"]