FROM golang:1.17-alpine as base

WORKDIR /app

RUN apk update &&\
    apk add --virtual build-dependencies build-base gcc

FROM base as build

COPY . .

RUN go mod download &&\
    go build -o /app/app

FROM alpine

WORKDIR /app

COPY --from=build /app/app ./

ENTRYPOINT [ "/app/app" ]