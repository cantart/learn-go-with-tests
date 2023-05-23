package main

import (
	"learngowithtests/mocking"
	"os"
)

func main() {
	mocking.Countdown(os.Stdout)
}
