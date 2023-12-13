// challenges/testing/begin/main_test.go
package main

import (
	"testing"
)

// write a test for letterCounter.count
func TestLetterCount(t *testing.T) {
	var letterTests = []struct {
		input string
		want  int
	}{
		{"#00", 0},
		{"a2_32_^_&/", 1},
		{"812_%6ab//", 2},
	}
	lc := letterCounter{}
	
	for _, tc := range letterTests {
		t.Run(tc.input, func(t *testing.T) {
			if got := lc.count(tc.input); got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

// write a test for numberCounter.count
func TestNumberCount(t *testing.T) {
	var numberTests = []struct {
		input string
		want  int
	}{
		{"#00", 2},
		{"abc_1,?/", 1},
		{"abc_12&B_^", 2},
	}
	nc := numberCounter{}

	for _, tc := range numberTests {
		t.Run(tc.input, func(t *testing.T) {
			if got := nc.count(tc.input); got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}

// write a test for symbolCounter.count
func TestSymbolCount(t *testing.T) {
	var symbolTests = []struct {
		input string
		want  int
	}{
		{"#00", 1},
		{"abc_1,?/", 4},
		{"abc_12&B_^", 4},
	}
	sc := symbolCounter{}

	for _, tc := range symbolTests {
		t.Run(tc.input, func(t *testing.T) {
			if got := sc.count(tc.input); got != tc.want {
				t.Errorf("got %v, wanted %v", got, tc.want)
			}
		})
	}
}
