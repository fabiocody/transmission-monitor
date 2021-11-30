package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Torrent struct {
    Hash string `gorm:"primaryKey"`
    Name string
}

var db *gorm.DB

func setupDB() {
    var err error
    db, err = gorm.Open(sqlite.Open(environment.databaseFile))
    handleErr(err)
    err = db.AutoMigrate(&Torrent{})
    handleErr(err)
}
