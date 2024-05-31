package main

import (
	"fmt"
	"math/rand"
	"mmp5/utils"
	"strconv"
	"time"
)

type classificator struct {
	vector []float64
	class  int
}

func main() {

	classificators := []classificator{}

	numbers, err := utils.DataReader("./data/data.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	var kString string

	fmt.Println("Give how many clusters you want to get (must be greater than 1): ")
	_, err = fmt.Scan(&kString)
	if err != nil {
		fmt.Println(err)
	}

	k, err := strconv.Atoi(kString)

	if err != nil {
		fmt.Println(err)
		return
	}

	if k > len(numbers) {
		k = len(numbers)
	}

	for i := 0; i < k; i++ {
		classificators = append(classificators, classificator{
			vector: numbers[i],
			class:  i,
		})
	}

	for i := k; i < len(numbers); i++ {
		rand.Seed(time.Now().UnixNano())
		classificators = append(classificators, classificator{
			vector: numbers[i],
			class:  rand.Intn(k),
		})
	}

	for i, v := range classificators {
    var arr [k]float64
	}

}
