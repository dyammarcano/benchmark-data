# Build stage
FROM golang:latest AS build

WORKDIR /app

COPY . .

RUN go build -o main .

# Final stage
FROM ubuntu:latest

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8080

CMD ["/app/main"]
