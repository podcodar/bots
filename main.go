package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"podcodar-discord-bot/services"
	"syscall"
)

func main() {
	// CLI args
	showScoreboard := flag.Bool("scoreboard", false, "Send scoreboard using daily channel")

	flag.Parse()

	switch {
	case *showScoreboard:
		fmt.Println("Sending scoreboard in daily channel...")
		services.SendScoreboard()

	default:
		// Start listening to messages
		err := services.ListenDailyMessages()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Wait here until CTRL-C or other term signal is received.
		fmt.Println("Bot is now running.  Press CTRL-C to exit.")
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sc

		// Cleanly close down the Discord session.
		services.CloseBot()
	}
}
