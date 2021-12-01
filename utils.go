package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func handleErr(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}

func getEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		logrus.Fatalf("Could not find environment variable %s\n", key)
	}
	return v
}

type Environment struct {
	debug                bool
	databaseFile         string
	transmissionHost     string
	transmissionUsername string
	transmissionPassword string
	transmissionHttps    bool
	transmissionPort     uint16
	telegramApiToken     string
	telegramChatId       int64
}

var environment Environment

func loadEnv() {
	dbFile, hasDbFile := os.LookupEnv("DATABASE_FILE")
	if !hasDbFile {
		dbFile = "transmission-monitor.db"
	}
	port, err := strconv.ParseUint(getEnv("TRANSMISSION_PORT"), 10, 16)
	handleErr(err)
	chatId, err := strconv.ParseInt(getEnv("TELEGRAM_CHAT_ID"), 10, 64)
	handleErr(err)
	environment = Environment{
		debug:                os.Getenv("DEBUG") == "1",
		databaseFile:         dbFile,
		transmissionHost:     getEnv("TRANSMISSION_HOST"),
		transmissionUsername: getEnv("TRANSMISSION_USERNAME"),
		transmissionPassword: getEnv("TRANSMISSION_PASSWORD"),
		transmissionHttps:    os.Getenv("TRANSMISSION_HTTPS") == "1",
		transmissionPort:     uint16(port),
		telegramApiToken:     getEnv("TELEGRAM_API_TOKEN"),
		telegramChatId:       chatId,
	}
}
