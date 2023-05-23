package mocking

import (
	"fmt"
	"io"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

type Sleeper interface {
	Sleep()
}

/*
This chapter will focus on It's an important skill to be able to slice up requirements as small as you can so you can have working software.
And, use DI to make function can work on test and terminal
*/

// refactoring some magic values into named constants.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

/*
* we want to control sleeping because slow tests ruin developer productivity
* If we can mock time.Sleep we can use dependency injection to use it instead of a "real" time.Sleep and then we can `spy on the calls`` to make assertions on them.

 */
