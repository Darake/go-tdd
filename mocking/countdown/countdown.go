package countdown

import (
	"fmt"
	"io"
)

const finalWord = "Go!"
const countdownStart = 3

func Start(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
