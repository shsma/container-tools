# Start from a Golang 1.20 base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY ../.. .

# Build the Go app inside the container
RUN go build -o main .

# Expose the port that the app listens on
EXPOSE 9003

# Run the app when the container starts
CMD ["./main"]

