# Build stage
FROM golang:1.20 AS build
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./main -mod=vendor ./cmd/myapp/main.go 

# Production stage
FROM debian:11
COPY --from=build /app/main /usr/bin/main
WORKDIR /usr/bin
EXPOSE 8080 
CMD ["./main"]
