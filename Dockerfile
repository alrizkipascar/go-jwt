FROM golang:1.21.6


COPY ./ /go/src/

WORKDIR /go/src/cmd

RUN go build -o ./bin/go-jwt

CMD ["./bin/go-jwt"]