# Use the official Golang image to build the Go application
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code and build the application
COPY . .
RUN go build -o /service1

# Debug: List the contents of the /app directory
RUN ls -l /app

# Expose port 9090
EXPOSE 9090

# Run the service
CMD ["/service1"]