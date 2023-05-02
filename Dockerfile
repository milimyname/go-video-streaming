# Dockerfile for go production
# Path: Dockerfile
# Compare this snippet from Dockerfile.dev:
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 3000
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
