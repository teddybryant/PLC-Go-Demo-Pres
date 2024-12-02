package main

import (
	"fmt"
)

func sumOfSquares(nums []int, ch chan int) {
	localSum := 0
	for _, num := range nums {
		localSum += num * num
	}
	ch <- localSum
}

func main() {
	numbers := []int{1, 2, 3, 4}
	ch := make(chan int)
	go sumOfSquares(numbers[:2], ch)
	go sumOfSquares(numbers[2:], ch)
	result1 := <-ch
	result2 := <-ch
	total := result1 + result2
	fmt.Println("Sum of squares:", total)
	fmt.Println("Hello World!")
}
