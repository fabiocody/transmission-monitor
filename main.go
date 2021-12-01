package main

import (
	"github.com/go-co-op/gocron"
	"time"
)

func main() {
	loadEnv()
	setupDB()
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(1).Minutes().Do(func() {
		transmission := getTransmissionClient()
		torrents := getTorrents(transmission)
		processTorrents(torrents)
	})
	handleErr(err)
	scheduler.StartBlocking()
}
