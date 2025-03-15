package sync

import (
	"sync"
	"testing"
)

/*
! Requirements
- Can invoke Inc method to count up its value on Counter
- Must be safe to use in a concurrent environment
*/

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				defer wg.Done()
				counter.Inc()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

// ! By waiting for wg.Wait() to finish before making our assertions we can be sure all of our goroutines have attempted to Inc the Counter.
// ! A Mutex must not be copied after first use. so use pass it to other func by pointer.
