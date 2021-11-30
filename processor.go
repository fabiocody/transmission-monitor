package main

import (
    "fmt"
    "github.com/hekmon/transmissionrpc/v2"
)

func processTorrents(torrents *[]transmissionrpc.Torrent) {
    torrentsMap := make(map[string]*string, len(*torrents))
    for _, t := range *torrents {
        if *t.PercentDone < 1.0 {
            var savedTorrent Torrent
            db.FirstOrCreate(&savedTorrent, Torrent{
                Hash: *t.HashString,
                Name: *t.Name,
            })
            torrentsMap[*t.HashString] = t.Name
            if environment.debug {
                fmt.Printf("Saving %#v", savedTorrent)
            }
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
