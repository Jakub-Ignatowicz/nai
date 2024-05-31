package main

import (
	"container/heap"
	"mpp6/utils"
)

func countLetters(sentence string) map[rune]int {
	letterCount := make(map[rune]int)

	// Use range to iterate over each rune in the string
	for _, char := range sentence {
		letterCount[char]++
	}

	return letterCount
}

func main() {
	letterCounts := countLetters("ala ma kota")

	pq := make(utils.PriorityQueue, len(letterCounts))

	i := 0
	for key, value := range letterCounts {
		pq[i] = utils.Item{
			Name:   key,
			Weight: value,
		}
		i++
	}
	heap.Init(&pq)
}
