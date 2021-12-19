package main

import (
	log "github.com/sirupsen/logrus"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
