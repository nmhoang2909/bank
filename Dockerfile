# build stage
FROM golang:1.22-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/main .
COPY dev.env .

EXPOSE 8080
CMD [ "/app/main" ]
