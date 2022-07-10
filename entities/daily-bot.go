package entities

type CreateDailyRecord struct {
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type DailyRecord struct {
	*CreateDailyRecord
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
}

type CreateDailyScoreboard struct {
	*CreateDailyRecord
	Points        int `json:"points"`
	CurrentStreak int `json:"currentStreak"`
}

type DailyScoreboard struct {
	*CreateDailyScoreboard
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
}
