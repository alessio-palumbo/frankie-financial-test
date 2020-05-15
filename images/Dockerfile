FROM golang:1.14.2-alpine3.11 as builder

WORKDIR /app

RUN apk update && \
  apk add git ca-certificates && \
  rm -rf /var/cache/apk/*

COPY go.mod .
COPY go.sum .
COPY main.go .
COPY pkg .

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main *.go

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/main /app/main

EXPOSE 8080

ENTRYPOINT [ "app/main" ]