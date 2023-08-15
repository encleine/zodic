package twit

import (
	"context"
	"encoding/json"
	"net/http"

	tscraper "github.com/n0madic/twitter-scraper"
)

type Scraper struct {
	scraper     *tscraper.Scraper
	username    string
	password    string
	openAccount bool
}

func NewScraper(username, password string) *Scraper {
	S := &Scraper{
		tscraper.New(),
		username, password,
		false,
	}
	S.login()
	return S
}

func (s *Scraper) login() {
	scraper := s.scraper
	if _0 := scraper.LoginOpenAccount(); _0 == nil {
		s.openAccount = true
		return
	}
	s.loginWithAccount()
}
func (s *Scraper) loginWithAccount() {
	if _0 := s.scraper.Login(s.username, s.password); _0 != nil {
		panic(_0)
	}
	s.openAccount = false
}

func (s *Scraper) getCookies() ([]byte, error) {
	return json.Marshal(s.scraper.GetCookies())
}

func (s *Scraper) loadCookie(js []byte) {
	var cookies []*http.Cookie
	json.Unmarshal(js, &cookies)
	s.scraper.SetCookies(cookies)
}

func (s *Scraper) GetTweets(username string, tweetAmount int) chan *tscraper.TweetResult {
	ch := make(chan *tscraper.TweetResult)
	scraper := s.scraper
	go func() {
		defer close(ch)
		for tweet := range scraper.GetTweets(context.Background(), username, tweetAmount) {
			if tweet.Error == nil {
				ch <- tweet
			} else if s.openAccount {
				s.loginWithAccount()
			} else {
				panic("your account got blocked")
			}
		}
	}()

	return ch
}
