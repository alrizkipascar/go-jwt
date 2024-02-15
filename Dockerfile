FROM golang:1.21.6

WORKDIR /usr/src/app

COPY ./* ./

RUN go build -o ./bin/RESTAPI

CMD ["./bin/RESTAPI"]