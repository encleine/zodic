package main

import (
	"fmt"
	"os"
	"time"

	"github.com/encleine/zodiac/api/bot"
	twit "github.com/encleine/zodiac/api/twit"
)

func main() {
	var hourAfter time.Time
	scraper := twit.NewScraper(
		os.Getenv("username"),
		os.Getenv("password"),
	)
	bot := bot.NewBot(os.Getenv("token"), "")

	for {
		for tweet := range scraper.GetTweets("hourIyhoroscope", 1) {
			fmt.Println(tweet.Text)
			bot.SendMessage(1806936826, tweet.Text)
			hourAfter = tweet.TimeParsed.Add(time.Hour)
		}
		if waitTime := hourAfter.Sub(time.Now()); waitTime > 0 {
			time.Sleep(waitTime)
		}
	}
}
