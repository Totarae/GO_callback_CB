# Use the official Golang image to build the Go application
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code and build the application
COPY . .
RUN go build -o /service2

# Run the service
CMD ["/app/service2"]
