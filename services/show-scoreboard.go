package services

import (
	"podcodar-discord-bot/repository"
)

func SendScoreboard() {
	content := repository.CreateScoreboardContent()

	// send message to daily channel
	ds := repository.MakeDiscordSession(config.DiscordToken)
	ds.Start()
	ds.SendMessage(config.DailyChannelId, content)
	ds.Close()
}
