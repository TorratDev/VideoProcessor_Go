# Build stage
FROM golang:1.23 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o video-processor main.go

# Final stage (runtime image)
FROM gcr.io/distroless/base-debian11

# Copy the built executable from the builder stage
COPY --from=builder /app/video-processor /video-processor

# Expose the necessary port (e.g., 8080 if using HTTP)
EXPOSE 8080

# Command to run the application
CMD ["/video-processor"]
