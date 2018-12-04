package fp

import (
	"fmt"
	"testing"
)

func TestPartial1(t *testing.T) {
	cases := []struct {
		in1, in2, expected string
	}{
		{"Hello","world", "Helloworldfred"},
		{"Hello","世界", "Hello世界fred"},
		{"", "", "fred"},
	}
	for _, c := range cases {
		var f = Partial1(p, c.in1)
		var actual = f(c.in2)
		if actual != c.expected {
			t.Errorf("Partial1(%q) == %q, want %q", c.in1, actual, c.expected)
		}
		fmt.Println(actual)
	}
}

func p(s1 string, s2 string) string {
	return s1 + s2 + "fred"
}

