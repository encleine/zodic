package bot

import tbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Telebot struct {
	bot       *tbot.BotAPI
	callbacks []Callback
	updates   tbot.UpdatesChannel
}
type Callback struct {
	Event    string
	Callback Callfunc
}
type Callfunc func(update tbot.Update)
