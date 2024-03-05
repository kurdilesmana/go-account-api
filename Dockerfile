FROM golang:1.20

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]
