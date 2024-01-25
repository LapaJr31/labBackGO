# Dockerfile
FROM golang:1.21.6-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /expense-tracker

EXPOSE 8080

CMD [ "/expense-tracker" ]
