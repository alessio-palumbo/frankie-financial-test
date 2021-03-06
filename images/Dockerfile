FROM golang:1.14.2-alpine3.11 as builder

WORKDIR /app

RUN apk update && \
  apk add git ca-certificates && \
  rm -rf /var/cache/apk/*

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go mod verify

COPY main.go .
COPY pkg ./pkg
COPY docs ./docs

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/main

EXPOSE 8080

ENTRYPOINT [ "app/main" ]