package main

import (
    "context"
    "github.com/hekmon/transmissionrpc/v2"
)

func getTransmissionClient() *transmissionrpc.Client {
    host := environment.transmissionHost
    username := environment.transmissionUsername
    password := environment.transmissionPassword
    config := transmissionrpc.AdvancedConfig{
        HTTPS: environment.transmissionHttps,
        Port:  environment.transmissionPort,
    }
    transmission, err := transmissionrpc.New(host, username, password, &config)
    handleErr(err)
    return transmission
}

func getTorrents(transmission *transmissionrpc.Client) *[]transmissionrpc.Torrent {
    torrents, err := transmission.TorrentGetAll(context.TODO())
    handleErr(err)
    return &torrents
}
