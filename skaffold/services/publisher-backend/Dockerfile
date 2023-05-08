FROM golang:1.20

WORKDIR /app

COPY . .
RUN go build -o main .
EXPOSE 9004

# Run  app when the container starts
CMD ["./main"]


