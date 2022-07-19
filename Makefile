run: main.go
	go run main.go

fmt:
	go fmt ./...

scoreboard: main.go
	go run main.go --scoreboard

build-image:
	docker build -t masouzajr/podcodar-discord-bot .

push-image: build-image
	docker push masouzajr/podcodar-discord-bot
