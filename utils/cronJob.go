package utils

import (
	"fmt"

	"github.com/robfig/cron"
)

type MyFuncType func()

func Schedule(spec string, callback MyFuncType) *cron.Cron {
	c := cron.New()
	c.AddFunc(spec, callback)
	c.Start()
	// Funcs may also be added to a running Cron
	c.AddFunc("@daily", func() { fmt.Println("Every day") })
	return c
}
