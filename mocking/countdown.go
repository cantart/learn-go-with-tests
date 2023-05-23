package mocking

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

/*
This chapter will focus on It's an important skill to be able to slice up requirements as small as you can so you can have working software.
And, use DI to make function can work on test and terminal
*/

// refactoring some magic values into named constants.
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}
