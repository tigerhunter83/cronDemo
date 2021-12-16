package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

const OneSecond = 1*time.Second + 50*time.Millisecond

func main() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() { fmt.Println("Every minute on sharp") })
	c.Start()
	select {}
}
