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

var Length string
var Source string
func init() {
	rootCmd.AddCommand(readCommand)
	rootCmd.AddCommand(syncCommand)
	readCommand.Flags().StringVarP(&Source, "source", "s", "", "source to read from")
	readCommand.Flags().StringVarP(&Length, "length", "l", "0", "no. of items to read")
}

func Execute() error {
	return rootCmd.Execute()
}
