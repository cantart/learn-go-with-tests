package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// ! Prefer to create a constructor which shows readers of your API that it would be better to not initialise the type yourself.

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock() // The zero value for a Mutex is an unlocked mutex.
	defer c.mu.Unlock()
	c.value++

}

func (c *Counter) Value() int {
	return c.value
}

/*
! Note:
- Use channels when passing ownership of data
- Use mutexes for managing state
- go vet
- Don't use structs embedding because it's convenient (embedding other struct to struct)
  by using it we can access the base’s fields directly that means you have exposed your API as public
*/
