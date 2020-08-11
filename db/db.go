package db

import (
	"fmt"
	"log"
	"time"

	"github.com/asdine/storm"
)

type Config struct {
	ID     int `storm:"id,increment"`
	Date   time.Time
	Energy int
}

var db *storm.DB

func Init(dbPath string) error {
	var err error
	db, err = storm.Open(dbPath)
	errFatal(err)
	return err
}

func ReadDB() []Config {
	var configs []Config
	err := db.All(&configs)
	errFatal(err)
	return configs
}

func AddLog(energy int) error {
	config := Config{Energy: energy, Date: time.Now()}
	err := db.Save(&config)
	if err != nil {
		return fmt.Errorf("Could not save config, %v", err)
	}
	return nil
}

func errFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
