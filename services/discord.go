package services

import (
	"log"
	"podcodar-discord-bot/repository"
	"podcodar-discord-bot/settings"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var config = settings.LoadConfig()
var session *discordgo.Session
var logger = log.Default()

func ListenDailyMessages() error {
	session, err := repository.MakeDiscordSession(config.DiscordToken)
	if err != nil {
		logger.Fatal("Error creating discord session", err)
	}

	// Add message handler
	session.AddHandler(dailyMessagesHandler)

	return session.Open()
}

func CloseBot() {
	session.Close()
}

func dailyMessagesHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	isDailyChannel := m.ChannelID == config.DailyChannelId
	if isDailyChannel {
		content := strings.ToLower(m.Content)
		dailyIdentifiers := []string{"o que eu fiz", "o que vou fazer"}
		isDailyMessage := containsAll(content, dailyIdentifiers)

		if isDailyMessage {
			logger.Println("✅ isDaily\n", m.Content)
			ComputereDailyScoreboard(m.Message)
		}
	}
}

func containsAll(text string, words []string) bool {
	for _, word := range words {
		if !strings.Contains(text, word) {
			return false
		}
	}
	return true
}
