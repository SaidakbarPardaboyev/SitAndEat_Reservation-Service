# Stage 1: Build stage
FROM golang:1.22.4 AS builder

WORKDIR /app

# Copy and download dependencies
COPY . ./
RUN go mod download

COPY .env ./

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp .

# Stage 2: Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/app.log .
COPY --from=builder /app/.env .

EXPOSE 6666

# Command to run the executable
CMD ["./myapp"]