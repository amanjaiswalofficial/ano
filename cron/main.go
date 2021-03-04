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
	API := api.InitAPI()
	
	Scheduler := job.GetScheduler()
	Scheduler.SubmitJob(job.PrintHelloWorld, 3)

	Scheduler.Start()
	API.Start(8080)

}
