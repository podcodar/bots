package repository

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type DiscordSession struct {
	Session *discordgo.Session
	token   string
}

func MakeDiscordSession(token string) *DiscordSession {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		panic(err)
	}

	return &DiscordSession{
		Session: session,
		token:   token,
	}
}

func (ds *DiscordSession) SendMessage(channel, content string) {
	_, err := ds.Session.ChannelMessageSend(channel, content)
	if err != nil {
		panic(err)
	}

	fmt.Println(content)
}

func (ds *DiscordSession) Start() {
	ds.Session.Open()
}

func (ds *DiscordSession) Close() {
	ds.Session.Close()
}
