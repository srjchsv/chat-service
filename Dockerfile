# Build stage
FROM golang:1.20 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./main ./cmd/myapp/main.go 

# Production stage
FROM alpine:3.15
COPY --from=build /app/main /usr/bin/main
WORKDIR /usr/bin
EXPOSE 8080 
CMD ["./main"]
