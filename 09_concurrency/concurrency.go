package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	// time.Sleep(2 * time.Second)
	return results
}

/*
## Concurrent
- goroutines
- channel: use to help organize and control the communication between the different processes
- race detector: test your code with race detector from `go test`
- select: use to synchronise processes

! you cannot directly add data to a channel outside of a goroutine. Channels are primarily used for communication and synchronization between goroutines. To add data to a channel, you need to do so within a goroutine.
*/
