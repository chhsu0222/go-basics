package saying

import "fmt"

// Greet returns a greet based on the input string s.
func Greet(s string) string {
	return fmt.Sprint("Welcome my dear ", s)
}
