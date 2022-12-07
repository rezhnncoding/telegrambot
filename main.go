package main

import (
	"fmt"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbot.NewBotAPI("5987539523:AAEHN7veaIwCuEUr_1Tiwf3HSpYdLWXQxCc")
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("bot connected successfully")
	u := tgbot.NewUpdate(0)
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbot.NewMessage(update.Message.Chat.ID, "reza")
		bot.Send(msg)
	}

}
