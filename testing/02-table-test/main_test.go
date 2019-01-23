package main

import "testing"

type testpair struct {
	data   []int
	answer int
}

func TestMysum(t *testing.T) {

	tests := []testpair{
		testpair{[]int{20, 22}, 42},
		testpair{[]int{13, -2}, 11},
		testpair{[]int{15, 10}, 25},
		testpair{[]int{7, 2}, 9},
	}

	for _, test := range tests {
		v := mySum(test.data...)
		if v != test.answer {
			t.Error("Expected", test.answer, "got", v)
		}
	}

}
