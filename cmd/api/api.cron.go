package main

import (
	c "github.com/robfig/cron"
	"github.com/tax/cmd/internal"
	"github.com/tax/lib/panics"
)

var (
	cronObj *c.Cron
)

func initCron(panicWrapper panics.Panics, ucase *internal.Usecase) {

	// Initialize cron object
	cronObj = c.New()

	//---------------------
	// START - Cron list
	//---------------------

	// -- Every 10 SECONDS
	cronObj.AddFunc("@every 10s", panicWrapper.Cron(func() {
		ucase.System.LogOpenFile()
	}))

	//---------------------
	// END - Cron list
	//---------------------

	cronObj.Start()
}
