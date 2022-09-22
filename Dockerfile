# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /server

COPY go.* ./
RUN go mod download

COPY main.go ./
COPY api/* ./api/
RUN go build -o /docker-my-server

EXPOSE 3001

CMD [ "/docker-my-server" ]