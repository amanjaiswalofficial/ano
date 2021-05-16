package cmd

import (
	"ano/job"
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var readCommand = &cobra.Command{
	Use:   "read",
	Short: "Fetch news from several RSS feeds",
	Long:  "Lists all the news from sources bases on flags",
	Run: func(cmd *cobra.Command, args []string) {

		sourceName, _ := cmd.Flags().GetString("source")
		length, _ := cmd.Flags().GetString("length")

		JSONData := job.ReadDataFromFile()

		itemCount, err := strconv.Atoi(length)
		if err != nil {
			itemCount = -1
		}
		PrintToCLI(JSONData, itemCount, sourceName)
	},
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "Sync news with urls",
	Long:  "Update the datastore with latest news from listed sources",
	Run: func(cmd *cobra.Command, args []string) {
		var urls []string

		file, err := os.Open("sources.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		job.SyncRSSFeeds(urls)
	},
}
