// testing/table-driven/begin/main_test.go
package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		name string
		input []int
		want int
	} {
		{"one", []int{1}, 1},
		{"two", []int{1,-3}, -2},
		{"three", []int{1,2,3}, 7},
	}

	// range over the tests and run them as subtests
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T){
			got := sum(tc.input...)
			if got != tc.want {
				t.Errorf("got %v, wanted %v",got, tc.want)
			}
		})
	}
}
