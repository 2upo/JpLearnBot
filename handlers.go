package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleDefault(bot *tgbotapi.BotAPI, update tgbotapi.Update) Status {
	status := Default

	if update.Message == nil {
		return status
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Command() {
	case "help":
		msg.Text = "I understand: \n/noteList\n/status\n/noteAdd"
	case "noteList":
		msg.Text = ReadFileLineByLine("./Notes.txt")
	case "status":
		msg.Text = "I'm ok."
	case "noteAdd":
		msg.Text = update.Message.From.UserName
		status = AddingNote
	default:
		msg.Text = "I understand: \n/noteList\n/status\n/noteAdd"
	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return status
}
