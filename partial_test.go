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

func TestCurry(t *testing.T) {
	cases := []struct {
		in1, in2, expected string
	}{
		{"Hello","world", "Helloworld"},
		{"Hello","世界", "Hello世界"},
		{"", "", ""},
	}
	for _, c := range cases {

		var add Add = h
		var curriedH  = add.Curry(c.in1)
		var curriedActual = curriedH(c.in2)
		if curriedActual != c.expected {
			t.Errorf("Partial1(%q) == %q, want %q", c.in1, curriedActual, c.expected)
		}
		fmt.Println(curriedActual)
	}
}

func h(s string, h string) string {
	return s + h
}
