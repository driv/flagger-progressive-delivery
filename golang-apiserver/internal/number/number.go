package number

import "time"

// Next takes a number as input and returns the next number.
// You can replace this logic with your own as needed.
func Next(current int) int {
	return current + 1
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	prev, curr := 0, 1
	for i := 1; i < n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
	// return Fibonacci(n-1) + Fibonacci(n-2)
}

func SlowFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Sleep(n int) int {
	time.Sleep(2 * time.Second)
	return n
}
