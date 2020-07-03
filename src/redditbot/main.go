package main

import (
	"fmt"
	"log"
	"os"

	"github.com/turnage/graw/reddit"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("ERROR: Gettin WD %s", err)
	}

	// TODO: Implementation should change here and we should retrieve our secrets from secretsmanager

	// cfg := BotConfig{
	// 	Agent: AGENT
	// 	App: App{
	// 	  ID:     ID,
	// 	  Secret: SECRET,
	// 	  Username: USERNAME,
	// 	  Password: PASSWORD,
	// 	}
	//   }

	// bot, err := reddit.NewBot(cfg)

	bot, err := reddit.NewBotFromAgentFile(fmt.Sprintf("%s/secrets.agent", path), 0)
	if err != nil {
		log.Fatalf("ERROR: Retrieving Agent File: %s", err)
	}

	submission, err := bot.GetPostSelf("LUC_team", "TEST GOLANG POST", "posting from bot")
	if err != nil {
		log.Fatalf("ERROR: Error while submitting the post -- %s", err)
	}

	fmt.Printf("%v", submission)
}
