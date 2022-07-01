# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.18-alpine AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o /api

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /api /api

ENV SERVER_PORT=3000

EXPOSE $SERVER_PORT

USER nonroot:nonroot

ENTRYPOINT ["/api"]
