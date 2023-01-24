// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCounter(t *testing.T) {
	tests := []struct{
		name string
		input string
		want int
	}{
		{input: "#00", want: 0},
		{input: "a2_32_^_&/)", want: 1},
		{input: "812_%6ab//", want: 2},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := letterCounter{}.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCounter(t *testing.T) {
	tests := []struct{
		name string
		input string
		want int
	}{
		{input: "#00", want: 2},
		{input: "abc_1,?/", want: 1},
		{input: "abc_12&8_^", want: 3},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := numberCounter{}.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}



// write a test for symbolCounter.count
func TestSymbolCounter(t *testing.T) {
	tests := []struct{
		name string
		input string
		want int
	}{
		{input: "#00", want: 1},
		{input: "abc_1,?/", want: 4},
		{input: "abc_12&8_^_", want: 5},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got := symbolCounter{}.count(tc.input)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}