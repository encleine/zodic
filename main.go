package main

import (
	"fmt"
	"os"
	"time"

	"github.com/encleine/zdic/bot"
	"github.com/encleine/zdic/twit"
)

func main() {
	var hourAfter time.Time
	scraper := twit.NewScraper(
		os.Getenv("username"),
		os.Getenv("password"),
	)
	tb := bot.NewBot(os.Getenv("token"))

  go func () {
    for {
      for tweet := range scraper.GetTweets("hourIyhoroscope", 1) {
        fmt.Println(tweet.Text)
        tb.SendToChannel("@ventoid", tweet.Text)
        hourAfter = tweet.TimeParsed.Add(time.Hour)
      }
      if waitTime := hourAfter.Sub(time.Now()); waitTime > 0 {
        time.Sleep(waitTime)
      }
    }
  }
}

