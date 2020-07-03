package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aldarisbm/redditbot/redditbot/secrets"
	"github.com/turnage/graw/reddit"
)

func main() {
	environment := os.Getenv("ENV")

	var bot reddit.Bot
	// If DEV we retrieve from agent file, if PROD we retrieve from SecretsManager
	if environment == "DEV" {
		path, err := os.Getwd()
		if err != nil {
			log.Fatalf("ERROR: Gettin WD -- %s", err)
		}

		bot, err = reddit.NewBotFromAgentFile(fmt.Sprintf("%s/redditbot/secrets.agent", path), 0)
		if err != nil {
			log.Fatalf("FATAL: Retrieving Agent File -- %s", err)
		}
	}
	if environment == "PROD" {
		secret, err := secrets.GetSecret()
		if err != nil {
			log.Fatalf("FATAL: Error while retrieving secret - %s", err)
		}
		cfg := reddit.BotConfig{
			Agent: secret.UserAgent,
			App: reddit.App{
				ID:       secret.ClientID,
				Secret:   secret.ClientSecret,
				Username: secret.Username,
				Password: secret.Password,
			},
		}

		bot, err = reddit.NewBot(cfg)
		if err != nil {
			log.Fatalf("FATAL: Error while creating bot from cfg - %s", err)
		}
	}

	submission, err := bot.GetPostSelf("LUC_team", "TEST GOLANG POST", "posting from bot")
	if err != nil {
		log.Fatalf("FATAL: Error while submitting the post -- %s", err)
	}

	fmt.Printf("%v", submission)

}
