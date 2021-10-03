package main

func main() {
}

type WebsiteChecker func(string) bool

type result struct {
	site string
	success bool
}

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {

	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.site] = r.success
	}


	return results
}
