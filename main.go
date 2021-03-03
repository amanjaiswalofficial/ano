package main

import (
	"time"

	"github.com/amanjaiswalofficial/ano/src"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
)

func main() {
	r := gin.Default()
	r.GET("/feed", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"N":       src.N,
		})
	})

	// Cron Scheduler setup
	sched := gocron.NewScheduler(time.UTC)
	sched.Every(3).Seconds().Do(src.Count)

	// Start the scheduler without blocking the current goroutine
	sched.StartAsync()

	r.Run(":8080")
}
