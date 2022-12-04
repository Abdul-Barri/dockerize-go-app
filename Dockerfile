FROM golang:1.16-alpine

MAINTAINER Abdul-Barri Lawal <lawalabdulbarri@gmail.com>

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o golang-app .

EXPOSE 8080

CMD [ "/app/golang-app" ]
