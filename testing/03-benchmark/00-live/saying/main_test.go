package saying

import "testing"

func TestGreet(t *testing.T) {
	s := Greet("John")
	if s != "Welcome my dear John" {
		t.Error("Expected", "Welcome my dear John", "Got", s)
	}
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
