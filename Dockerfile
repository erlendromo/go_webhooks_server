# Build stage
FROM golang:1.21.5 AS builder

# Maintainers and authors
LABEL maintainer="Cloudenberg"
LABEL authors="Erlend, Arthur, Oskar, Martin"

# Workdir name for image
WORKDIR /webhooks

# Copy the entire project directory
COPY . .

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o executable ./cmd/main.go

# Expose internal port
EXPOSE 8081

# Run executable binary
ENTRYPOINT ["./executable"]
