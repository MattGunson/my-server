# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY api/* ./api/
COPY *.go ./

RUN go build -o /docker-my-server

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-my-server /docker-my-server

EXPOSE 3001

USER nonroot:nonroot

ENTRYPOINT ["/docker-my-server"]