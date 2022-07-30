package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

var wg sync.WaitGroup

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	chanel := make(chan FreqMap)
	m := make(FreqMap)
	defer close(chanel)

	wg.Add(len(l))

	for i := 0; i < len(l); i++ {
		i := i
		go func() {
			defer wg.Done()
			chanel <- Frequency(l[i])
		}()
	}
	go func() {
		wg.Wait()
	}()
	for freq := range chanel {
		for i, f := range freq {
			m[i] += f
		}
	}
	return m
}
