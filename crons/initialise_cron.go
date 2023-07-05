package crons

import (
	"github.com/robfig/cron"
)

func Initialise() error {
	c := cron.New()

	// Add your cron job to the scheduler
	c.AddFunc("0 * * * * *", CronJob1) // Run the job every minute

	// Start the cron scheduler in a Goroutine
	go c.Start()

	return nil
}
