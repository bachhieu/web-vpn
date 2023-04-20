package utils

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type MyFuncType func()

func Schedule(spec string, callback MyFuncType) *cron.Cron {
	fmt.Print("cron start!!!!!!!!!!!!!!")
	var c = cron.New()
	c.AddFunc(spec, callback)
	c.Start()
	return c
}

func AddFunc(c *cron.Cron, spec string, callback MyFuncType) *cron.EntryID {
	id, _ := c.AddFunc(spec, callback)
	return &id
}

func RemoveFunc(c *cron.Cron, id cron.EntryID) {
	c.Remove(id)
}
