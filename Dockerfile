FROM golang:1.25-alpine AS build

WORKDIR /build

COPY ./go.mod .

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:latest AS topz

WORKDIR /app

COPY --from=build /build/main ./main

ENV PORT=8000

ENTRYPOINT ["/app/main"]