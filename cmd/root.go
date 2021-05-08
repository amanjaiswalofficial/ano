package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	rootCmd = &cobra.Command{
		Use:   "ano",
		Short: "ano is an awesome RSS reader with joyful colors",
	}
)

func init() {
	rootCmd.AddCommand(readCommand)
	rootCmd.AddCommand(syncCommand)
}

func Execute() error {
	return rootCmd.Execute()
}
