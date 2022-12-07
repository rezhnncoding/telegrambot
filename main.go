package main

import (
	"fmt"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
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
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if len(update.Message.Text) == 0 {
			bot.DeleteMessage(tgbot.DeleteMessageConfig{update.Message.Chat.ID, update.Message.MessageID})
		}
		msg := tgbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)

	}
}
