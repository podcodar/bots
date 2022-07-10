# build project in a two-step container
FROM denoland/deno:latest AS builder
WORKDIR /bots/
COPY . ./
RUN deno compile -c deno.jsonc --allow-read --allow-env --allow-net --import-map=./imports.json --no-check -o podcodar-bot mod.ts

FROM denoland/deno:latest
WORKDIR /app/
COPY --from=builder /bots/podcodar-bot ./
CMD ["./podcodar-bot"]

