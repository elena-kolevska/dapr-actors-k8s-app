# Start from the official Golang base image
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the Working Directory inside the container
COPY *.go ./

# Build the Go app
RUN go build -o main .
RUN CGO_ENABLED=0 GOOS=linux go build -o /actors-service

# Start a new stage from scratch
FROM alpine:latest  

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /actors-service /actors-service

ENV PORT 3333
ENV NS ""
# Expose port 3333 to the outside world
EXPOSE ${PORT}

# Command to run the executable
CMD ["/actors-service"]
