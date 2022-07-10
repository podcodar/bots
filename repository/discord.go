package repository

import (
	"github.com/bwmarrin/discordgo"
)

func MakeDiscordSession(token string) (*discordgo.Session, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return session, nil
}
