## Build Stage
FROM golang:1.18-alpine as go-builder

WORKDIR /usr/src/app

COPY . .

RUN apk add -u -t build-tools curl git && \ 
    go build -o main *.go && \ 
    apk del build-tools && \ 
    rm -rf /var/cache/apk/*

## Runtime Stage
FROM alpine:latest

WORKDIR /usr/src/app

RUN apk --no-cache add ca-certificates

COPY --from=go-builder /usr/src/app/main /usr/src/app/main

EXPOSE 8080

ENTRYPOINT [ "/usr/src/app/main" ]