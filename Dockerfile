# Use an official Golang runtime as a parent image
FROM golang:1.20-alpine

#  设置国内源
ENV GOPROXY=https://goproxy.cn,direct

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download -x

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o /app/main .

WORKDIR /app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]