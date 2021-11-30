package main

import (
    "fmt"
    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendNotification(torrent *Torrent) {
    bot, err := tgbotapi.NewBotAPI(environment.telegramApiToken)
    handleErr(err)
    msgString := fmt.Sprintf("%s has finished downloading\n", torrent.Name)
    if environment.debug {
        fmt.Println(msgString)
    }
    msg := tgbotapi.NewMessage(environment.telegramChatId, msgString)
    _, err = bot.Send(msg)
    handleErr(err)
}
