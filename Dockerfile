FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o telegram-catbot .

WORKDIR /dist
RUN cp /build/telegram-catbot .

# Build a small image
FROM scratch
COPY --from=builder /dist/telegram-catbot /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/telegram-catbot"]
