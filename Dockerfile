FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

# Build the Go app
RUN go build -o /app/main .

WORKDIR /app

EXPOSE 3000

CMD ["./main"]