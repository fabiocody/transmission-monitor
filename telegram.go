package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func sendNotification(torrent *Torrent) {
	bot, err := tgbotapi.NewBotAPI(environment.telegramApiToken)
	handleErr(err)
	msgString := fmt.Sprintf("%s has finished downloading\n", torrent.Name)
	logrus.Debugln(msgString)
	msg := tgbotapi.NewMessage(environment.telegramChatId, msgString)
	_, err = bot.Send(msg)
	handleErr(err)
}
