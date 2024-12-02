package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func ParallelSum(numbers []int, routines int) (int, error) {
	if routines <= 0 {
		return 0, errors.New("number of routines must be greater than zero")
	}

	var wg sync.WaitGroup
	sum := 0
	var mu sync.Mutex

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

	chunkSize := len(numbers) / routines
	for i := 0; i < routines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == routines-1 {
			end = len(numbers)
		}
		wg.Add(1)
		go worker(numbers[start:end])
	}

	wg.Wait()
	return sum, nil
}

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

type Rectangle struct {
	Width, Height float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	defer fmt.Println("\nDeferred: clean up resources")
	fmt.Println("### Go Demo ###")
	fmt.Println("\n# Section 1: Basic Syntax and Typing")

	var explicitInt int = 100
	inferredInt := 100

	fmt.Printf("explicitInt: %d (type: %T)\n", explicitInt, explicitInt)
	fmt.Printf("inferredInt: %d (type: %T)\n", inferredInt, inferredInt)

	defer fmt.Println("Printed last")
	fmt.Println("Rprinted first")

	//panic("Something is wrong")

	var x int = 10

	fmt.Printf("Variable x initialized with value %d\n", x)
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	fmt.Println("For loop demonstration:")
	for i := 0; i < 3; i++ {
		fmt.Printf("Loop iteration %d\n", i)
	}
	fmt.Println("Conditional for loop:")
	i := 5
	for i > 0 {
		fmt.Println("Countdown:", i)
		i--
	}

	var y = 20
	fmt.Printf("Variable y inferred as type %T with value %d\n", y, y)
	var s Shape = Rectangle{Width: 5, Height: 10}
	fmt.Printf("Area of rectangle: %.2f\n", s.Area())

	num := 42
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	day := "Tuesday"
	fmt.Printf("Day: %s\n", day)
	switch day {
	case "Monday":
		fmt.Println("1st Day")
	case "Tuesday":
		fmt.Println("2nd Day")
	default:
		fmt.Println("3rd Day or more")
	}

	fmt.Println("\n# Section 2: Error Handling")
	fmt.Println("SafeDivide function:")
	divResult, err := SafeDivide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result of 10 / 2: %d\n", divResult)
	}

	divResult, err = SafeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result of 10 / 0: %d\n", divResult)
	}

	fmt.Println("\n# Section 3: Parallelism and Concurrency")

	filename := "numbers.txt"
	fmt.Println("Reading numbers from file...")
	numbers, err := ReadIntegersFromFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	start := time.Now()
	sum, err := ParallelSum(numbers, 4)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Parallel sum of squares: %d\n", sum)
	fmt.Printf("Time taken: %s\n", elapsed)

	fmt.Println("\n### End of Demo ###")
}
