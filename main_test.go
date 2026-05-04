package main

import "testing"

func TestSumOfTwoNumber(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}
	for _, tt := range tests {
		got := SumOfTwoNumber(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("SumOfTwoNumber(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSubtractTwoNumbers(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-1, -1, 0},
		{100, 200, -100},
	}
	for _, tt := range tests {
		got := SubtractTwoNumbers(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("SubtractTwoNumbers(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMultiplyTwoNumbers(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{3, 4, 12},
		{0, 100, 0},
		{-2, 5, -10},
		{7, 7, 49},
	}
	for _, tt := range tests {
		got := MultiplyTwoNumbers(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("MultiplyTwoNumbers(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
