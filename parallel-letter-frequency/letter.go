package letter

import "sync"

// ConcurrentFrequency finds the frequency of letters in an array of strings.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	counts := make(chan FreqMap) // channel to send counts for each text
	totals := FreqMap{}

	// map: get freq counts for each text concurrently
	for _, text := range texts {
		wg.Add(1)
		go func(text string) {
			defer wg.Done()
			counts <- Frequency(text)
		}(text)
	}

	// reduce: tally up the counts as they're sent on the channel
	go func() {
		for c := range counts {
			for k, v := range c {
				totals[k] += v
			}
		}
	}()

	wg.Wait()
	return totals
}
