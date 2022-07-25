package repository

import (
	"encoding/json"
	"fmt"
	"podcodar-discord-bot/entities"
	"podcodar-discord-bot/settings"
	"strings"
	"time"

	pgo "github.com/supabase/postgrest-go"
)

var (
	client  *pgo.Client
	config  = settings.LoadConfig()
	headers = map[string]string{}
	schema  = "public"
)

func NewClient() *pgo.Client {
	if client != nil {
		return client
	}

	client = pgo.
		NewClient(config.SupabaseUrl, schema, headers).
		TokenAuth(config.SupabaseToken)

	if client.ClientError != nil {
		panic(client.ClientError)
	}

	return client
}

func FindUserScoreByUserId(userId string) *entities.DailyScoreboard {
	if client == nil {
		client = NewClient()
	}

	queryBuilder := client.
		From("daily_scoreboard").
		Select("*", "exact", false).
		Eq("userId", userId)

	data, count, err := queryBuilder.Single().Execute()
	if err != nil {
		if count == 0 {
			return nil
		}
		panic(err)
	}

	var dailyScoreboard entities.DailyScoreboard
	err = json.Unmarshal(data, &dailyScoreboard)
	if err != nil {
		panic(err)
	}

	return &dailyScoreboard
}

func AddDailyRecord(dailyRecord entities.CreateDailyRecord) *entities.DailyRecord {
	if client == nil {
		client = NewClient()
	}

	data, _, err := client.
		From("daily_record").
		Insert(dailyRecord, true, "", "", "exact").
		Single().
		Execute()

	if err != nil {
		panic(err)
	}

	var dailyScoreboard entities.DailyRecord
	err = json.Unmarshal(data, &dailyScoreboard)
	if err != nil {
		panic(err)
	}

	return &dailyScoreboard
}

func AddDailyScoreboard(dailyRecord entities.CreateDailyRecord) *entities.DailyScoreboard {
	if client == nil {
		client = NewClient()
	}

	createDailyScoreboard := entities.CreateDailyScoreboard{
		Points:            1,
		CurrentStreak:     1,
		CreateDailyRecord: &dailyRecord,
	}

	data, _, err := client.
		From("daily_scoreboard").
		Insert(createDailyScoreboard, true, "", "", "exact").
		Single().
		Execute()

	if err != nil {
		panic(err)
	}

	var dailyScoreboard entities.DailyScoreboard
	err = json.Unmarshal(data, &dailyScoreboard)
	if err != nil {
		panic(err)
	}

	return &dailyScoreboard
}

func SaveDailyScoreboard(dailyRecordToUpdate *entities.DailyScoreboard) *entities.DailyScoreboard {
	if client == nil {
		client = NewClient()
	}

	data, _, err := client.
		From("daily_scoreboard").
		Upsert(dailyRecordToUpdate, "", "", "").
		Single().
		Execute()

	if err != nil {
		panic(err)
	}

	var dailyScoreboard entities.DailyScoreboard
	err = json.Unmarshal(data, &dailyScoreboard)
	if err != nil {
		panic(err)
	}

	return &dailyScoreboard
}

func CountUserActivityLastDays(userId string, days int) int {
	initialDate := time.Now().AddDate(0, 0, -days)

	if client == nil {
		client = NewClient()
	}

	queryBuilder := client.
		From("daily_record").
		Select("*", "exact", false).
		Gt("created_at", initialDate.Format("2006-01-02")).
		Eq("userId", userId)

	_, count, err := queryBuilder.Execute()
	if err != nil {
		panic(err)
	}

	return int(count)
}

func ScoreboardRanking(top int) []entities.DailyScoreboard {
	if client == nil {
		client = NewClient()
	}

	orderOptions := pgo.OrderOpts{Ascending: false}
	queryBuilder := client.
		From("daily_scoreboard").
		Select("*", "exact", false).
		Order("points", &orderOptions).
		Limit(top, "")

	ranking := []entities.DailyScoreboard{}
	_, err := queryBuilder.ExecuteTo(&ranking)
	if err != nil {
		panic(err)
	}

	return ranking
}

var MAX_NAME_LENGTH = 25

func CreateScoreboardContent() string {
	now := time.Now()
	content := fmt.Sprintf("\n**Scoreboard - %s**\n\n", now.Format("02.01.2006"))
	codeQuote := "```\n"
	content += codeQuote

	// get ranked users from scoreboard repository
	for index, user := range ScoreboardRanking(10) {
		nameSize := len(user.Name)
		spacesString := strings.Repeat(" ", MAX_NAME_LENGTH-nameSize)
		name := fmt.Sprintf("%s: %s", user.Name, spacesString)

		content += fmt.Sprintf("%d. %s | %d\n", index, name, user.Points)
	}
	content += codeQuote

	return content
}
