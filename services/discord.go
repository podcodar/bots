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
	ds := repository.MakeDiscordSession(config.DiscordToken)
	session = ds.Session

	// Add message handler
	session.AddHandler(dailyMessagesHandler)
	session.AddHandler(commandsHandler)

	return session.Open()
}

func AddCommands() {
	log.Println("Adding commands...")

	for index, cmd := range commands {
		command, err := session.ApplicationCommandCreate(session.State.User.ID, config.GuildId, cmd)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", cmd.Name, err)
		}
		registeredCommands[index] = command
	}
}

func CloseBot() {
	log.Println("Removing commands...")
	for _, v := range registeredCommands {
		err := session.ApplicationCommandDelete(session.State.User.ID, config.GuildId, v.ID)

		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
	session.Close()
}

var registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "scoreboard",
		Description: "Mostra o ranking do dia",
	},
}

var commandsMapper = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"scoreboard": scoreboardCommand,
}

func commandsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if handler, ok := commandsMapper[i.ApplicationCommandData().Name]; ok {
		handler(s, i)
	}
}

func scoreboardCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: repository.CreateScoreboardContent(),
		},
	})
}

func dailyMessagesHandler(_ *discordgo.Session, m *discordgo.MessageCreate) {
	isDailyChannel := m.ChannelID == config.DailyChannelId
	if isDailyChannel {
		content := strings.ToLower(m.Content)
		dailyIdentifiers := []string{"o que eu fiz", "o que vou fazer"}
		isDailyMessage := containsAll(content, dailyIdentifiers)

		if isDailyMessage {
			logger.Println("✅ isDaily\n", m.Content)
			ComputeDailyScoreboard(m.Message)
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
