package main

import (
	"fmt"
	"log"

	"github.com/aldarisbm/redditbot/redditbot/secrets"
)

func main() {
	secret, err := secrets.GetSecret()

	if err != nil {
		log.Fatal("SOFATA")
	}

	cfg := BotConfig{
		Agent: secret.UserAgent
		App: App{
		  ID:     secret.ClientID,
		  Secret: secret.ClientSecret,
		  Username: secret.Username,
		  Password: secret.Password,
		}
	  }

	bot, err := reddit.NewBot(cfg)s

	// submission, err := bot.GetPostSelf("LUC_team", "TEST GOLANG POST", "posting from bot")
	// if err != nil {
	// 	log.Fatalf("ERROR: Error while submitting the post -- %s", err)
	// }

	// fmt.Printf("%v", submission)
}
