# 1. Build project
FROM arm64v8/golang:1.18 as builder

WORKDIR /app/
COPY ./ ./

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o podcodar-discord-bot .

# 2. Pack compiled code
FROM arm64v8/alpine:latest as packer

WORKDIR /root/
COPY --from=builder /app/podcodar-discord-bot ./podcodar-discord-bot

RUN apk --no-cache add ca-certificates

CMD ["/root/podcodar-discord-bot"]
