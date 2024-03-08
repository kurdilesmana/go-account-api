FROM golang:1.20-alpine

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/server apps/gl/cmd/main.go
CMD ["./bin/server"]
