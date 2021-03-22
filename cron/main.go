package main

import (
	"cron/api"
	"cron/job"
)

func InitApp() int {
	// to be updated later
	return 0
}

func main() {

	job.SetupEnvironment("../config.json")

	var urls []string
	urls = append(urls, "https://nerdist.com/feed/")
	urls = append(urls, "https://edm.com/.rss/full/")

	API := api.InitAPI()

	Scheduler := job.GetScheduler()
	//Scheduler.SubmitJob(job.PrintHelloWorld, 3)

	job.SyncRSSFeeds(urls)

	Scheduler.Start()
	API.Start(8081)

}
