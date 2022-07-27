package letter

import "fmt"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

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
	defer close(chanel)
	var m FreqMap
	for i := 0; i < len(l); i++ {
		i := i
		go func() {
			chanel <- Frequency(l[i])
		}()
	}
	for freq := range chanel {
		fmt.Println(freq)
	}
	return m
}
