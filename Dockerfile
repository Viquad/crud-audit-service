# Build stage 
FROM golang:1.18-alpine as builder
WORKDIR /app
COPY cmd cmd
COPY internal internal
COPY pkg pkg
COPY go.mod go.sum ./
RUN go mod download
RUN go build -v -o /app/main ./cmd/main.go

# Run stage 
FROM alpine:latest
WORKDIR /app
COPY configs configs
COPY --from=builder /app/main .

CMD ["/app/main"]