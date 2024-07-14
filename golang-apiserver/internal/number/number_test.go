package number

import "testing"

func TestNext(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"next number for 1", 1, 2},
		{"next number for 42", 42, 43},
		{"next number for 100", 100, 101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Next(tt.input)
			if result != tt.expected {
				t.Errorf("NextNumber(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"Fibonacci of 0", 0, 0},
		{"Fibonacci of 1", 1, 1},
		{"Fibonacci of 2", 2, 1},
		{"Fibonacci of 3", 3, 2},
		{"Fibonacci of 4", 4, 3},
		{"Fibonacci of 5", 5, 5},
		{"Fibonacci of 6", 6, 8},
		{"Fibonacci of 7", 7, 13},
		{"Fibonacci of 10", 10, 55},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Fibonacci(tt.input)
			if result != tt.expected {
				t.Errorf("Fibonacci(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
