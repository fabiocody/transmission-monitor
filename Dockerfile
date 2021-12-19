# syntax=docker/dockerfile:1

## BUILD
FROM golang:1.17 AS build
WORKDIR /app
RUN mkdir /data
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o /transmission-monitor

## Deploy
FROM gcr.io/distroless/base
WORKDIR /
COPY --from=build /data /data
COPY --from=build /transmission-monitor /transmission-monitor
ENTRYPOINT ["/transmission-monitor"]
