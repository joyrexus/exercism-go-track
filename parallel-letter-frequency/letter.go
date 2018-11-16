package letter

// ConcurrentFrequency finds the frequency of letters in an array of strings.
func ConcurrentFrequency(texts []string) FreqMap {
	counts := make(chan FreqMap, len(texts))
	totals := FreqMap{}

	// map: get freq counts for each text concurrently
	for _, text := range texts {
		go func(text string) {
			counts <- Frequency(text)
		}(text)
	}

	// reduce: tally up the counts as they're sent on the channel
	for range texts {
		for k, v := range <-counts {
			totals[k] += v
		}
	}

	return totals
}
