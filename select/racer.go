package racer

import (
	"fmt"
	"net/http"
	"time"
)

/*
! NOTE
* chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool
* When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
*/

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	// fmt.Printf("url: %s, response: %T\n", url, <-ch)
	return ch
}
