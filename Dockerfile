### 
FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o main /app/cmd/arog-gora-quote-api

ENTRYPOINT ["/app/main"]


