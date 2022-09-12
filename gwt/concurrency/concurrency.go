package concurrency

import "time"

type WebsiteChacker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChacker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// recieve expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	time.Sleep(2 * time.Second)

	return results
}
