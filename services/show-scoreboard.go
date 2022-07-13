package services

import (
	"fmt"
	"podcodar-discord-bot/repository"
	"strings"
	"time"
)

var maxNameLength = 25

func SendScoreboard() {
	now := time.Now()
	content := fmt.Sprintf("\n**Scoreboard - %s**\n\n", now.Format("02.01.2006"))
	codeQuote := "```\n"
	content += codeQuote

	// get ranked users from scoreboard repository
	for index, user := range repository.ScoreboardRanking(10) {
		nameSize := len(user.Name)
		spacesString := strings.Repeat(" ", maxNameLength-nameSize)
		name := fmt.Sprintf("%s: %s", user.Name, spacesString)

		content += fmt.Sprintf("%d. %s | %d\n", index, name, user.Points)
	}
	content += codeQuote

	// send message to daily channel
	ds := repository.MakeDiscordSession(config.DiscordToken)
	ds.Start()
	ds.SendMessage(config.DailyChannelId, content)
	ds.Close()
}
