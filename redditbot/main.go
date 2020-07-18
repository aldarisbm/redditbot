package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aldarisbm/redditbot/redditbot/secrets"
	"github.com/turnage/graw/reddit"
)

var githubURL = "https://raw.githubusercontent.com/snori74/linuxupskillchallenge/master"

func main() {
	env := os.Getenv("ENV")

	bot, err := getBot(env)
	if err != nil {
		log.Fatalf("FATAL: Error while getting bot - %s", err)
	}
	day := getDay()
	text, err := getText(day)
	if err != nil {
		log.Fatalf("FATAL: Error while getting text - %s", err)
	}
	submission, err := bot.GetPostSelf("LUC_team", "TEST MARKDOWN POST", text)
	if err != nil {
		log.Fatalf("FATAL: Error while submitting the post - %v", err)
	}

	fmt.Printf("%v", submission)
}

func getBot(env string) (bot reddit.Bot, err error) {
	// If DEV we retrieve from agent file, if PROD we retrieve from SecretsManager
	switch env {
	case "DEV":
		path, err := os.Getwd()
		if err != nil {
			log.Fatalf("ERROR: Gettin WD - %s", err)
			return bot, err
		}
		bot, err = reddit.NewBotFromAgentFile(fmt.Sprintf("%s/secrets.agent", path), 0)
		if err != nil {
			log.Fatalf("FATAL: Retrieving Agent File - %s", err)
			return bot, err
		}
	case "PROD":
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
	default:
		return bot, fmt.Errorf("Only DEV or PROD accepted for $ENV")
	}
	return bot, nil
}

func getText(day int) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%d.md", githubURL, day))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func getDay() int {
	return 2
}
