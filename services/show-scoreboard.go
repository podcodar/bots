package services

import (
	"fmt"
	"podcodar-discord-bot/repository"
	"time"
)

func SendScoreboard() {
	now := time.Now()
	resultString := fmt.Sprintf("\n# Scoreboard - %s\n\n", now.Format("01.02.2006"))

	// get ranked users from scoreboard repository
	for index, user := range repository.ScoreboardRanking(10) {
		tabSeparator := "\t"
		if len(user.Name) <= 10 {
			tabSeparator += "\t"
		}
		resultString += fmt.Sprintf("%d. %s: %s %d\n", index+1, user.Name, tabSeparator, user.Points)
	}

	fmt.Println(resultString)
	// TODO: send message to daily channel
}
