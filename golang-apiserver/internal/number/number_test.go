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
