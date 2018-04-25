package main

import (
	"github.com/robfig/cron"
	"github.com/xzdbd/portal/internal/api"
)

func main() {
	api.SyncFileItems()
	s := cron.New()
	s.AddFunc("0 0 1 * * *", api.SyncFileItems)
	s.Start()

	r := api.GinAPIRouter()
	r.Run(":8080")
}
