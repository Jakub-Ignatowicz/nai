package main

import (
	"container/heap"
	"fmt"
	"math"
	"mpp6/utils"
)

func countLetters(sentence string) map[rune]int {
	letterCount := make(map[rune]int)

	for _, char := range sentence {
		letterCount[char]++
	}

	return letterCount
}

func createInitialPriorityQueue(sentence string) utils.PriorityQueue {
	letterCounts := countLetters(sentence)

	pq := make(utils.PriorityQueue, len(letterCounts))

	i := 0
	for key, value := range letterCounts {
		pq[i] = utils.Item{
			Name:   string(key),
			Weight: value,
		}
		i++
	}
	heap.Init(&pq)

	return pq
}

func main() {
	str := "hakuma tata ma"

	huffmanEncoding(str)

	str = "ala ma kota a kot ma ale"

	huffmanEncoding(str)

	str = "hello world"

	huffmanEncoding(str)

	str = "ashdfsdfh ,basmndfbas,ndfbamsfmnbs,fwaefhyuiyabwefauydsvcuysdvbdsaucisadc"

	huffmanEncoding(str)
}

func huffmanEncoding(str string) {
	fmt.Println(fmt.Sprintf("Starting string: %s", str))

	pq := createInitialPriorityQueue(str)

	bitSizeOfChar := 1

	for math.Pow(2, float64(bitSizeOfChar)) < float64(pq.Len()) {
		bitSizeOfChar++
	}

	preencodedSize := bitSizeOfChar * len(str)

	result := fmt.Sprintf("Total bit length of not encoded string is: %d", preencodedSize)

	fmt.Println(result)

	encodedChars := make(map[string]string)

	for pq.Len() > 1 {
		left := heap.Pop(&pq).(utils.Item)
		right := heap.Pop(&pq).(utils.Item)

		newName := left.Name + right.Name
		newWeight := left.Weight + right.Weight

		newItem := utils.Item{
			Name:   newName,
			Weight: newWeight,
		}

		heap.Push(&pq, newItem)

		for _, char := range left.Name {
			encodedChars[string(char)] = "0" + encodedChars[string(char)]
		}

		for _, char := range right.Name {
			encodedChars[string(char)] = "1" + encodedChars[string(char)]
		}
	}

	fmt.Println()
	for key, value := range encodedChars {
		fmt.Printf("Char: %s, Encoded: %s\n", key, value)
	}
	fmt.Println()

	lenghtAfterEncoding := 0

	pq2 := createInitialPriorityQueue(str)

	for pq2.Len() > 0 {
		item := heap.Pop(&pq2).(utils.Item)
		lenghtAfterEncoding += len(encodedChars[item.Name]) * item.Weight
	}

	result2 := fmt.Sprintf("After encoding we got %d bits in this string so we saved %d bits", lenghtAfterEncoding, preencodedSize-lenghtAfterEncoding)

	fmt.Println(result2)
	fmt.Println()
}
