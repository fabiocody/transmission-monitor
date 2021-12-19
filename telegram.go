package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

func SendNotification(torrent *Torrent) {
	bot, err := tgbotapi.NewBotAPI(Args.ApiToken)
	HandleErr(err)
	msgString := fmt.Sprintf("%s has finished downloading\n", torrent.Name)
	log.Debugln(msgString)
	msg := tgbotapi.NewMessage(Args.ChatId, msgString)
	_, err = bot.Send(msg)
	HandleErr(err)
}
