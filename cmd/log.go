package cmd

import (
	"fmt"
	"strconv"

	"github.com/nahumsa/HourlyTracker/db"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Logs your energy at a given time",
	Run: func(cmd *cobra.Command, args []string) {
		energy, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		err = db.AddLog(energy)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(logCmd)
}
