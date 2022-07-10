# build project in a two-step container
FROM denoland/deno:latest

WORKDIR /app/
COPY . ./

CMD ["deno", "run", "--allow-env", "--allow-read", "--allow-net", "mod.ts"]

