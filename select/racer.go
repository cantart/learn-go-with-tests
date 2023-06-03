package racer

import (
	"net/http"
)

/*
! NOTE
* chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool
* When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
*/

func Racer(a, b string) string {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
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
