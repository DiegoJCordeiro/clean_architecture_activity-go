FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o clean_architecture_activity ./cmd/clean_architecture_activity

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder --chmod=755 /app/clean_architecture_activity .
COPY app.env* ./
ENV MONGODB_URI=mongodb://admin:admin123@mongodb:27017
EXPOSE 8080
CMD ["./rater-limiter-activity"]