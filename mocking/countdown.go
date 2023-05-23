package mocking

import (
	"fmt"
	"io"
)

/*
This chapter will focus on It's an important skill to be able to slice up requirements as small as you can so you can have working software.
And, use DI to make function can work on test and terminal
*/
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}
