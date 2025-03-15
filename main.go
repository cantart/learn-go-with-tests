package main

import (
	mocking "learngowithtests/08_mocking"
	"os"
	"time"
)

func main() {
	// sleeper := &mocking.DefaultSleeper{}
	// mocking.Countdown(os.Stdout, sleeper)

	sleeper := &mocking.ConfigurableSleeper{}
	sleeper.SetDuration(5 * time.Second)
	sleeper.SetSleepFunc(time.Sleep)
	mocking.Countdown(os.Stdout, sleeper)
}
