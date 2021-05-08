package main

import (
	"ano/cmd"
	"ano/job"
)

func main() {
	var urls []string
	urls = append(urls, "https://nerdist.com/feed/")
	urls = append(urls, "https://edm.com/.rss/full/")

	
	job.SyncRSSFeeds(urls)
	cmd.Execute()
}
