# Use the official Go image as the base for building the binary
FROM golang:1.21.12 AS builder
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

# Use a lightweight image to run the binary
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .
COPY .env .

# Expose the port your Go app listens on
EXPOSE 7070

# Run the Go app
CMD ["./app"]
