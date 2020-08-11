package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/nahumsa/HourlyTracker/cmd"
	"github.com/nahumsa/HourlyTracker/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "energy.db")
	err := db.Init(dbPath)
	must(err)
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
