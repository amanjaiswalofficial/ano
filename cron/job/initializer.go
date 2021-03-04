package job

import (
	"time"

	"github.com/go-co-op/gocron"
)

type Scheduler struct{ obj *gocron.Scheduler }


func GetScheduler() Scheduler {
	// Cron Scheduler setup
	sched := gocron.NewScheduler(time.UTC)

	return Scheduler{sched}
}

// SubmitJob will submit a cron job repeating every duration no. of seconds
func (s *Scheduler) SubmitJob(j Job, duration int) {
	s.obj.Every(duration).Seconds().Do(j)
}

// Start will start all the cron jobs submitted previously
func (s *Scheduler) Start() {
	s.obj.StartAsync()
}
