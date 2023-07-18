package bot

import (
	"log"
	"net/http"
	"strings"

	tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBot(Bot_token, url string) *Telebot {
	var bot, _0 = tbot.NewBotAPI(Bot_token)
	if _0 != nil {
		log.Panic(_0)
	}
	tb := &Telebot{
		bot:       bot,
		callbacks: []Callback{},
	}
	// tb.webHook(url)
	return tb
}

func (tb *Telebot) SendMessage(chatId int64, message string) {
	msg := tbot.NewMessage(chatId, message)
	tb.bot.Send(msg)
}

func (tb *Telebot) SendToChannel(username, message string) {
	msg := tbot.NewMessageToChannel(username, message)
	tb.bot.Send(msg)
}

func (tb *Telebot) webHook(url string) {
	bot := tb.bot
	wh, _ := tbot.NewWebhook(url + bot.Token)

	_, _0 := bot.Request(wh)
	if _0 != nil {
		log.Fatal(_0)
	}

	tb.updates = bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":8443", nil)
}

func (tb *Telebot) GetUpdates() {
	u := tbot.NewUpdate(0)
	u.Timeout = 60
	go func() {
		for update := range tb.updates {
			for _, callback := range tb.callbacks {
				switch event := strings.Split(callback.Event, ":"); {
				case event[0] == "message" && update.Message != nil:
					if len(event) == 2 && event[1] != update.Message.From.UserName {
						return
					}
					callback.Callback(update)
				case event[0] == "channel post" && update.ChannelPost != nil:
					if len(event) == 2 && event[1] != update.ChannelPost.From.UserName { // channel post : channel username
						return
					}
					callback.Callback(update)
				case event[0] == "inlineQuery" && update.InlineQuery != nil:
					callback.Callback(update)
				case event[0] == "chosen InlineResult" && update.ChosenInlineResult != nil:
					callback.Callback(update)
				case event[0] == "callbackQuery" && update.CallbackQuery != nil:
					callback.Callback(update)
				default:
					log.Panic("unknown event")
				}
			}
		}
	}()
}

func (tb *Telebot) on(event string, callback Callfunc) {
	tb.callbacks = append(tb.callbacks, Callback{event, callback})
}
