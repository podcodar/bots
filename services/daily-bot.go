package services

import (
	"podcodar-discord-bot/entities"
	"podcodar-discord-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func ComputeDailyScoreboard(msg *discordgo.Message) {
	dailyRecord := entities.CreateDailyRecord{
		UserId: msg.Author.ID,
		Name:   msg.Author.Username,
	}

	dailyScoreboard := repository.FindUserScoreByUserId(dailyRecord.UserId)
	if dailyScoreboard == nil {
		repository.AddDailyScoreboard(dailyRecord)
	} else {
		dailyScoreboard.Points++
		// calculate extra points
		if repository.CountUserActivityLastDays(dailyRecord.UserId, 7) >= 5 {
			dailyScoreboard.Points++
		}

		dailyScoreboard.CurrentStreak++
		// if no report yesterday, reset current streak
		if repository.CountUserActivityLastDays(dailyRecord.UserId, 1) == 0 {
			dailyScoreboard.CurrentStreak = 0
		}

		repository.SaveDailyScoreboard(dailyScoreboard)
	}
	repository.AddDailyRecord(dailyRecord)
}
