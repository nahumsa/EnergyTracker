package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "TaskManager",
	Short: "This is a CLI Energy Tracker",
}
