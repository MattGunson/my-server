# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY api/* ./api/
COPY db/* ./db/
COPY *.go ./

RUN go build -o /docker-my-server

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

ENV PGHOST containers-us-west-33.railway.app
ENV PGPORT 6744
ENV PGUSER postgres
ENV PGPASSWORD AY0FNCBK456XYmYkDCIP
ENV PGDATABASE railway
ENV DATABASE_URL postgresql://${PGUSER}:${PGPASSWORD}@${PGHOST}:${PGPORT}/${PGDATABASE}

COPY --from=build /docker-my-server /docker-my-server

EXPOSE 3001

USER nonroot:nonroot

ENTRYPOINT ["/docker-my-server"]