package main

import (
	"mocking/countdown"
	"os"
)

func main() {
	defaultSleeper := &countdown.DefaultSleeper{}
	countdown.Start(os.Stdout, defaultSleeper)
}
