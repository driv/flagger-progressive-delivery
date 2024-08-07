package number

import "time"

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
}

func BadFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Sleep(n int) int {
	time.Sleep(2 * time.Second)
	return n
}
