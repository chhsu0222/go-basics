package mymath

import "testing"

func TestCenteredAvg(t *testing.T) {

	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{1, 4, 6, 8, 100}, 6},
		test{[]int{0, 8, 10, 1000}, 9},
		test{[]int{9000, 4, 10, 8, 6, 12}, 9},
		test{[]int{123, 744, 140, 200}, 170},
	}

	for _, v := range tests {
		avg := CenteredAvg(v.data)
		if v.answer != avg {
			t.Error("Expected", v.answer, "Got", avg)
		}
	}
}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{9000, 4, 10, 8, 6, 12}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}
