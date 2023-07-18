package main

import (
	"fmt"
	"os"
	"time"

	"github.com/encleine/ventoid/bot"
	"github.com/encleine/ventoid/twit"
)

func main() {
	var hourAfter time.Time
	scraper := twit.NewScraper(
		os.Getenv("username"),
		os.Getenv("password"),
	)
	tb := bot.NewBot(os.Getenv("token"), os.Getenv("hookurl"))

	for {
		for tweet := range scraper.GetTweets("hourIyhoroscope", 1) {
			fmt.Println(tweet.Text)
			tb.SendToChannel("ventoid", tweet.Text)
			hourAfter = tweet.TimeParsed.Add(time.Hour)
		}
		if waitTime := hourAfter.Sub(time.Now()); waitTime > 0 {
			time.Sleep(waitTime)
		}
	}
}
