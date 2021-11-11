package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "taskmanager",
	Short: "taskmanager is a CLI task manager",
}
