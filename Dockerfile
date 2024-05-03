# We specify the base image we need for our GO app
FROM golang:1.22.2 as builder

# Create /app directory within the image to hold our application source code
WORKDIR /app

# we copy everything in the root directory into our /app directory
COPY go.mod .
COPY go.sum .

# Install dependencies
RUN go mod download

# Copy source files in app directory
COPY . .
### # Copy top-level Go files
### COPY *.go ./          
### # Copy services package files
### COPY services/ ./services/    
### # Copy cmd package files
### COPY cmd/ ./cmd/ 
### # Copy config package files
### COPY config/ ./config/ 
### # Copy common package files
### COPY common/ ./common/ 


# Build the app with optional configuration
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/go-docker-multistage
FROM alpine:latest
COPY --from=builder /opt/go-docker-multistage /opt/go-docker-multistage

# tells docker that the container listens on specified network ports at runtime
EXPOSE 8080

# command to be used to execute when the image is used to start the container
ENTRYPOINT [ "/opt/go-docker-multistage" ]