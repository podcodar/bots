# 1. Build project
FROM golang:1.18 as builder

WORKDIR /app/
COPY ./ ./

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o podcodar-discord-bot .

# 2. Pack compiled code
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/podcodar-discord-bot ./podcodar-discord-bot

RUN apk --no-cache add ca-certificates

CMD ["./podcodar-discord-bot"]
