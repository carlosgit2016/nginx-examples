FROM golang:1.19.11-bullseye as build

WORKDIR /app
COPY . .

RUN go build main.go

FROM alpine:latest

WORKDIR app
COPY --from=build /app/main main

ENTRYPOINT ["/bin/sh", "-c", "/app/main"]
