package cmd

import (
	"ano/job"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func PrintToCLI(rssData job.RSSData, itemCount int, sourceName string) {

	red := color.New(color.FgHiRed)
	cyan := color.New(color.FgCyan, color.Bold)
	yellow := color.New(color.FgHiYellow, color.Underline)
	magenta := color.New(color.FgHiMagenta)

	matched := false
	for urls := range rssData {
		if strings.Contains(urls, sourceName){
			sourceName = urls
			matched = true
		}
	}

	if !matched {
		red.Println("Unable to find source, please use 'ano list sources' for valid source names ")
		return
	}

	if itemCount == -1 || itemCount > 50 {
		itemCount = 50
	}

	for i := 0; i < itemCount; i++ {
		item := rssData[sourceName][i]

		timeObject, _ := time.Parse(time.RFC1123, item.Date)
		timeString := strconv.Itoa(timeObject.Day()) + " " + timeObject.Month().String() + " " + strconv.Itoa(timeObject.Year())

		cyan.Printf("%02d.", i+1)
		magenta.Print("[", timeString, "] ")
		cyan.Print(item.Title, "\n")
		yellow.Print(item.Link, "\n")

	}
}
