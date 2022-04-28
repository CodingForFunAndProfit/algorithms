package main

import (
	"fmt"
	"time"

	"github.com/CodingForFunAndProfit/algorithms/sorting"
)

func main() {
	arr := []int{5, 2, 6, 9, 3, 1, 4, 7, 8, 0}
	fmt.Println("Original: ", arr)
	start := time.Now()
	arr = sorting.Bubblesort(arr)
	elapsed := time.Since(start)
	fmt.Println("Took", elapsed.Nanoseconds())
	fmt.Println("Sorted: ", arr)
}
