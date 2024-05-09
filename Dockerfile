
# Use the official Go image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the application source code
COPY . .

# Download and install Go dependencies
RUN go mod download

# # Build the Go application
# RUN CGO_ENABLED=0 GOOS=linux go build -o mytodoapp
RUN go build -o idpservice cmd/main.go

# # Use a lightweight Alpine image as a final base image
# FROM alpine:latest

# # Set the working directory inside the container
# WORKDIR /app

# # Copy the built binary from the builder stage
# COPY --from=builder /app/idpservice .

# Expose the port the application listens on
EXPOSE 3000

# Command to run the executable
CMD ["./idpservice"]
