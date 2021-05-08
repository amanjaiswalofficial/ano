package cmd

import (
	"fmt"
	"ano/job"

	"github.com/fatih/color"
)

func PrintToCLI(rssData job.RSSData) {
	for url, data := range rssData {
		color.Yellow(url)
		for count, newsItem := range data {
			headline := fmt.Sprint(count+1, ". [", newsItem.Date, "] ", newsItem.Title, "\n", newsItem.Link)
			color.Cyan(headline)
		}
	}
}
