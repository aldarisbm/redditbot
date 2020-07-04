package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aldarisbm/redditbot/redditbot/secrets"
	"github.com/turnage/graw/reddit"
)

func main() {
	env := os.Getenv("ENV")

	bot, err := getBot(env)
	if err != nil {
		log.Fatalf("FATAL: Error while getting bot - %s", err)
	}

	submission, err := bot.GetPostSelf("LUC_team", "TEST GOLANG POST", "posting from bot")
	if err != nil {
		log.Fatalf("FATAL: Error while submitting the post - %s", err)
	}

	fmt.Printf("%v", submission)
}

func getBot(env string) (bot reddit.Bot, err error) {
	// If DEV we retrieve from agent file, if PROD we retrieve from SecretsManager
	if env == "DEV" {
		path, err := os.Getwd()
		if err != nil {
			log.Fatalf("ERROR: Gettin WD - %s", err)
			return bot, err
		}
		bot, err = reddit.NewBotFromAgentFile(fmt.Sprintf("%s/redditbot/secrets.agent", path), 0)
		if err != nil {
			log.Fatalf("FATAL: Retrieving Agent File - %s", err)
			return bot, err
		}
	}
	if env == "PROD" {
		secret, err := secrets.GetSecret()
		if err != nil {
			log.Fatalf("FATAL: Error while retrieving secret - %s", err)
			return bot, err
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
			return bot, err
		}
	}
	return bot, nil
}
