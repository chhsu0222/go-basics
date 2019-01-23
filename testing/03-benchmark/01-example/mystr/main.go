package mystr

import "strings"

// Cat is a self-defined function to concatenate the strings in a slice.
func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// Join uses the built-in strings.Join function.
func Join(xs []string) string {
	return strings.Join(xs, " ")
}
