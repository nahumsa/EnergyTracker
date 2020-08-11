package cmd

import (
	"fmt"

	"github.com/nahumsa/EnergyTracker/db"

	"github.com/spf13/cobra"
)

var histCmd = &cobra.Command{
	Use:   "history",
	Short: "Show all your logs.",
	Run: func(cmd *cobra.Command, args []string) {

		logs := db.ReadDB()

		if len(logs) == 0 {
			fmt.Println("You have no logs.")
			return
		}

		fmt.Println("You have the following logs:")
		for _, val := range logs {
			fmt.Printf("Time: %s \t Energy:%v\n", val.Date.Format("2006-01-02 15:04"), val.Energy)
		}
	},
}

func init() {
	RootCmd.AddCommand(histCmd)
}
