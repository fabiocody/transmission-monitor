package main

import (
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.SetReportCaller(true)
	loadEnv()
	if environment.debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	setupDB()
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Cron("*/1 * * * *").Do(checkTorrents)
	handleErr(err)
	logrus.Info("Started")
	scheduler.StartBlocking()
}

func checkTorrents() {
	logrus.Debug("Checking torrents")
	transmission := getTransmissionClient()
	torrents := getTorrents(transmission)
	torrentsMap := make(map[string]*string, len(*torrents))
	for _, t := range *torrents {
		if *t.PercentDone < 1.0 {
			var savedTorrent Torrent
			db.FirstOrCreate(&savedTorrent, Torrent{
				Hash: *t.HashString,
				Name: *t.Name,
			})
			torrentsMap[*t.HashString] = t.Name
			logrus.Debugf("Saving %#v", savedTorrent)
		}
	}
	var dbTorrents []Torrent
	db.Find(&dbTorrents)
	for _, t := range dbTorrents {
		if _, ok := torrentsMap[t.Hash]; !ok {
			db.Delete(&t)
			sendNotification(&t)
		}
	}
}
