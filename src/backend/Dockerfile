FROM golang:alpine AS builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy wait-for-postgres and add permission
COPY wait-for.sh .
RUN chmod +x wait-for.sh

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Test application
RUN go test

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine:3.11

COPY --from=builder /dist/main /
EXPOSE 8000

# /main will be executed in docker-compose.yml after wait-for.sh is done
