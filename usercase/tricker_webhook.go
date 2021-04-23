package usercase

import (
	"fmt"
	"github.com/robfig/cron"
)

func CronToTriggerWebhook()  {
	fmt.Println("start cron")
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		fmt.Println("Every hour thirty")
	})
	c.Start()

}
