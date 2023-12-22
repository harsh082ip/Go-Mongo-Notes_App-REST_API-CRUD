# Use the official Golang image with the latest version
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container at /app 
COPY . /app

# Build the Go application 
RUN go build -o app .

# Expose port 8083 to the outside world
EXPOSE 8083

# Command to run the executable
CMD ["./app"]
