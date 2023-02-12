FROM golang:1.17-alpine AS builder

WORKDIR /apps

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./app/app.go

FROM alpine:3.14

RUN apk update && apk --no-cache add curl ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /apps
COPY --from=builder /apps/api /apps/api

EXPOSE 8080
ENTRYPOINT ["/apps/api"]