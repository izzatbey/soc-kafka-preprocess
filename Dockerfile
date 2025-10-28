FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o preprocess-service ./cmd/consumer

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/preprocess-service /preprocess-service
ENTRYPOINT ["/preprocess-service"]
