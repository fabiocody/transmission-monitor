# syntax=docker/dockerfile:1

## BUILD
FROM golang:1.17 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /transmission-monitor
RUN mkdir /data

## Deploy
FROM gcr.io/distroless/base
WORKDIR /
COPY --from=build /transmission-monitor /transmission-monitor
COPY --from=build /data /data
ENTRYPOINT ["/transmission-monitor"]
