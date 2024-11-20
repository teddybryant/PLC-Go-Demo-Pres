package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ParallelSum calculates the sum of squares using goroutines
func ParallelSum(numbers []int) int {
	var wg sync.WaitGroup
	sum := 0
	var mu sync.Mutex

	// Worker function
	worker := func(nums []int) {
		defer wg.Done()
		localSum := 0
		for _, num := range nums {
			localSum += num * num
		}
		mu.Lock()
		sum += localSum
		mu.Unlock()
	}

	chunkSize := len(numbers) / 4
	for i := 0; i < 4; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == 3 {
			end = len(numbers) // Handle remainder
		}
		wg.Add(1)
		go worker(numbers[start:end])
	}

	wg.Wait()
	return sum
}

// ReadIntegersFromFile reads integers from a file
func ReadIntegersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, num)
		}
	}

	return numbers, scanner.Err()
}

func main() {
	filename := "numbers.txt" // Create a file with integers separated by spaces/newlines

	// Read numbers from file
	numbers, err := ReadIntegersFromFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Measure time for parallel computation
	start := time.Now()
	sum := ParallelSum(numbers)
	elapsed := time.Since(start)

	fmt.Printf("Parallel sum of squares: %d\n", sum)
	fmt.Printf("Time taken: %s\n", elapsed)
}
