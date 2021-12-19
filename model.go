package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Torrent struct {
	Hash string `gorm:"primaryKey"`
	Name string
}

var DB *gorm.DB

func SetupDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open(Args.Database))
	HandleErr(err)
	err = DB.AutoMigrate(&Torrent{})
	HandleErr(err)
}
