FROM golang:1.20

WORKDIR /app

COPY . .
RUN go build -o main .
EXPOSE 9004

CMD ["./main"]


