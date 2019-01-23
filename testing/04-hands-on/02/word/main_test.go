package word

import (
	"testing"

	"github.com/chhsu0222/go-basics/testing/04-hands-on/02/quote"
)

func TestCount(t *testing.T) {
	c := Count(" foo  bar yes ")
	if c != 3 {
		t.Error("Expected", 3, "Got", c)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount(" foo  bar yes foo ")
	for k, v := range m {
		switch k {
		case "foo":
			if v != 2 {
				t.Error("Expected", 2, "Got", v)
			}
		case "bar":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		case "yes":
			if v != 1 {
				t.Error("Expected", 1, "Got", v)
			}
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
