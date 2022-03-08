package main

import (
	"mocking/countdown"
	"os"
)

func main() {
	countdown.Start(os.Stdout)
}
