package main

import "testing"

func TestMysum(t *testing.T) {
	v := mySum(8, 5)
	if v != 13 {
		t.Error("Expected 13, got ", v)
	}
}
