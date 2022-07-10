package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Settings struct {
	DiscordToken   string
	DailyChannelId string
	SupabaseUrl    string
	SupabaseToken  string
}

var settings *Settings
var logger = log.Default()

func LoadConfig() Settings {
	if settings != nil {
		return *settings
	}

	logger.Println("Loading config...")

	err := godotenv.Load()
	if err != nil {
		logger.Println("No .env file found")
	}

	// fill up settings values
	settings = &Settings{
		DiscordToken:   os.Getenv("DISCORD_TOKEN"),
		DailyChannelId: os.Getenv("DAILY_CHANNEL"),
		SupabaseUrl:    os.Getenv("SUPABASE_URL"),
		SupabaseToken:  os.Getenv("SUPABASE_TOKEN"),
	}

	return *settings
}
