run: main.go
	go run main.go

build-image:
	docker build -t masouzajr/podcodar-discord-bot .

push-image: build-image
	docker push masouzajr/podcodar-discord-bot
