FROM golang:1.23.2 AS development
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex
EXPOSE 8000
CMD reflex -g '*.go' go run main.go --start-service