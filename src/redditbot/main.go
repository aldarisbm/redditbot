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

	bot, err := reddit.NewBotFromAgentFile(fmt.Sprintf("%s/secrets.agent", path), 0)
	if err != nil {
		log.Fatalf("ERROR: Retrieving Agent File: %s", err)
	}

	harvest, _ := bot.Listing("/r/golang", "")
	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}
}
