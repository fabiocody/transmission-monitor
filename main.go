package main

import (
	"github.com/alexflint/go-arg"
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	log.SetReportCaller(true)
	arg.MustParse(&Args)
	if Args.Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	SetupDB()
	if Args.Daemon {
		scheduler := gocron.NewScheduler(time.UTC)
		_, err := scheduler.Cron("*/1 * * * *").Do(CheckTorrents)
		HandleErr(err)
		log.Info("Started")
		scheduler.StartBlocking()
	} else {
		CheckTorrents()
	}
}

func CheckTorrents() {
	log.Debug("Checking torrents")
	transmission := GetTransmissionClient()
	torrents := GetTorrents(transmission)
	torrentsMap := make(map[string]*string, len(*torrents))
	for _, t := range *torrents {
		if *t.PercentDone < 1.0 {
			var savedTorrent Torrent
			DB.FirstOrCreate(&savedTorrent, Torrent{
				Hash: *t.HashString,
				Name: *t.Name,
			})
			torrentsMap[*t.HashString] = t.Name
			log.Debugf("Saving %#v", savedTorrent)
		}
	}
	var dbTorrents []Torrent
	DB.Find(&dbTorrents)
	for _, t := range dbTorrents {
		if _, ok := torrentsMap[t.Hash]; !ok {
			DB.Delete(&t)
			SendNotification(&t)
		}
	}
}
