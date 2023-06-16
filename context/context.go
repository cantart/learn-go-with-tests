package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// non-sense implementation
// func Server(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		store.Cancel()
// 		fmt.Fprint(w, store.Fetch())
// 	}
// }

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return // todo: log error however you like
		}
		fmt.Fprint(w, data)
	}
}

/*
context has a method Done() which returns a channel which gets sent a signal when the context is "done" or "cancelled".
We want to listen to that signal and call store.Cancel if we get it but we want to ignore it if our Store manages to Fetch before it.
*/
// context is goroutines
