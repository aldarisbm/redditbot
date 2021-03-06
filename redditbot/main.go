package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aldarisbm/graw/reddit"
	"github.com/aldarisbm/redditbot/redditbot/secrets"
)

var githubURL = "https://raw.githubusercontent.com/snori74/linuxupskillchallenge/master"

func main() {
	env := os.Getenv("ENV")

	bot, err := getBot(env)
	if err != nil {
		log.Fatalf("FATAL: Error while getting bot - %s", err)
	}

	postDay := getPostDay(time.Now())
	if postDay < 0 {
		log.Fatal("FATAL: NOOP")
	}

	textBody, err := getTextBody(postDay)
	if err != nil {
		log.Fatalf("FATAL: Error while getting text - %s", err)
	}

	submission, err := bot.GetPostSelf("luc_team", "TEST MARKDOWN POST", textBody)
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

func getTextBody(day int) (string, error) {
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

func getPostDay(now time.Time) int {
	dayInt := now.Weekday()
	todaysDate := now.Day()
	firstMondayOfMonth := getFirstMondayOfMonth(now, time.Now().Month())

	if todaysDate < firstMondayOfMonth {
		return -1
	}

	switch dayInt {
	case time.Monday:
		return -1
	case time.Tuesday:
		return -1
	case time.Wednesday:
		return -1
	case time.Thursday:
		return -1
	case time.Friday:
		return -1
	default:
		return 1
	}
}

func getFirstMondayOfMonth(now time.Time, month time.Month) (firstMonday int) {
	year, month, hour, minute, second, nanoSecond, location :=
		now.Year(),
		month,
		now.Hour(),
		now.Minute(),
		now.Second(),
		now.Nanosecond(),
		now.Location()

	for day := 1; day < 8; day++ {
		if time.Date(year, month, day, hour, minute, second, nanoSecond, location).Weekday() == time.Monday {
			firstMonday = day
		}
	}
	return firstMonday
}
