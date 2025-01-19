FROM golang:1.22.5-alpine3.20 AS build

WORKDIR /app

COPY main.go greeting.go ./

RUN go build -o /app/template-go main.go greeting.go

FROM alpine:3.20
LABEL org.opencontainers.image.source https://github.com/nu12/template-go

WORKDIR /app

COPY --from=build /app/template-go /app/template-go

ENTRYPOINT [ "/app/template-go" ]