package main

import (
	"mocking/countdown"
	"os"
	"time"
)

func main() {
	sleeper := &countdown.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	countdown.Start(os.Stdout, sleeper)
}
