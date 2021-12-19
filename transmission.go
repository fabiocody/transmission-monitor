package main

import (
	"context"
	"github.com/hekmon/transmissionrpc/v2"
)

func GetTransmissionClient() *transmissionrpc.Client {
	host := Args.Host
	username := Args.Username
	password := Args.Password
	config := transmissionrpc.AdvancedConfig{
		HTTPS: Args.Https,
		Port:  Args.Port,
	}
	transmission, err := transmissionrpc.New(host, username, password, &config)
	HandleErr(err)
	return transmission
}

func GetTorrents(transmission *transmissionrpc.Client) *[]transmissionrpc.Torrent {
	torrents, err := transmission.TorrentGetAll(context.TODO())
	HandleErr(err)
	return &torrents
}
