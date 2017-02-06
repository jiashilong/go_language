package main

import "fmt"

func getAverage(arr []int, n int) float32 {
	var sum int = 0
	var average float32 = 0.0

	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	average = float32(sum) / float32(n)
	return average
}

func main() {
	var arr = []int{1, 2, 0, 4}
	fmt.Println(getAverage(arr, len(arr)))
}
