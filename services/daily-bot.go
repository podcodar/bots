package services

import (
	"podcodar-discord-bot/entities"
	"podcodar-discord-bot/repository"
	"time"

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

		computeExtraPoints(dailyScoreboard)

		computeStreak(dailyRecord.UserId, dailyScoreboard)

		repository.SaveDailyScoreboard(dailyScoreboard)
	}
	repository.AddDailyRecord(dailyRecord)
}

var REPORT_STREAK = 4

func computeExtraPoints(dailyScoreboard *entities.DailyScoreboard) {
	// add extra points if user has 5 standups on the current week
	weekday := time.Now().Weekday()

	if weekday >= 4 {
		if repository.CountUserActivityLastDays(dailyScoreboard.UserId, int(weekday)) == REPORT_STREAK {
			dailyScoreboard.Points++
		}
	}
}

func computeStreak(userId string, dailyScoreboard *entities.DailyScoreboard) {
	dailyScoreboard.CurrentStreak++
	// if no report yesterday, reset current streak
	if repository.CountUserActivityLastDays(userId, 1) == 0 {
		dailyScoreboard.CurrentStreak = 0
	}
}
