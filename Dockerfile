#############################
# Use the official Golang image as the base image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module and Go sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/app ./cmd/app/main.go

#############################
# Use the official Alpine image as the base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root

# Copy from builder stage
COPY --from=builder /app/bin/app .
COPY --from=builder /app/config /root/config
COPY --from=builder /app/.env /root/.env

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Setting the CONFIG_PATH environment variable
ENV CONFIG_PATH=./config/prod.yaml

# Command to run the application
CMD ["./app"]
