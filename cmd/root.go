package cmd

import (
	"ano/job"

	"github.com/spf13/cobra"
	//"os"
	//homedir "github.com/mitchellh/go-homedir"
	//"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ano",
	Short: "Fetch news from several RSS Feeds",
	Long:  `ano simply prints all the recent news articles from RSS Feeds`,
	Run: func(cmd *cobra.Command, args []string) {
		JSONData := job.ReadDataFromFile()
		PrintToCLI(JSONData)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
