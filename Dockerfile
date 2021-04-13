# Pull down latest golang image
FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -v -u github.com/gorilla/mux

CMD ./main

EXPOSE 10000