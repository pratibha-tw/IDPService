# Start from a Golang base image
FROM golang:1.22.3-alpine3.19

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o idpservice cmd/main.go

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./idpservice"]