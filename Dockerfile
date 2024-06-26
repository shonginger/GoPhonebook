# We specify the base image we need for our GO app
FROM golang:1.22.2 as builder

# Create /app directory within the image to hold our application source code
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Copy source files from the project directory into the appropriate directories in the /app directory
COPY . .

# Build the app with optional configuration
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create a new image based on Alpine
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Set the entry point for the container
CMD [ "./main" ]