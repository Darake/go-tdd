package countdown

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}

const finalWord = "Go!"
const countdownStart = 3

func Start(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
