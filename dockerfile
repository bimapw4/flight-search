FROM golang:1.21-alpine

RUN apk add --no-cache git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build main-service
RUN go build -o flight-api ./main/main.go

# Build provider-service
RUN go build -o flight-provider ./provider/main.go

EXPOSE 3000

# NOTE: Command akan di-override di docker-compose
CMD ["/bin/sh"]
