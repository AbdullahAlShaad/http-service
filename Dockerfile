# Use the official Golang image for building the Go application
FROM golang:1.21-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules definition files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /http-go-service

# Use a minimal base image for the runtime environment
FROM alpine:latest

# Copy the built application from the builder stage
COPY --from=build /http-go-service /http-go-service

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application
CMD ["/http-go-service"]
