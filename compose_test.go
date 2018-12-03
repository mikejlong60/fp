package fp

import (
	"fmt"
	"testing"
)

func TestCompose(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"Hello, world", "Hello, worldfredsimon"},
		{"Hello, 世界", "Hello, 世界fredsimon"},
		{"","fredsimon"},
	}
	for _, c := range cases {

	    var f = Compose(f, g)
    	var actual = f(c.in)
    	if actual != c.expected {
        			t.Errorf("Compose(%q) == %q, want %q", c.in, actual, c.expected)
        		}
    	fmt.Println(actual)
	}
}

func TestAndThen(t *testing.T) {
	cases := []struct {
		in, expected string
	}{
		{"Hello, world", "Hello, worldfredsimon"},
		{"Hello, 世界", "Hello, 世界fredsimon"},
		{"","fredsimon"},
	}
	for _, c := range cases {

	    var f = AndThen(g, f)
    	var actual = f(c.in)
    	if actual != c.expected {
        			t.Errorf("Compose(%q) == %q, want %q", c.in, actual, c.expected)
        		}
    	fmt.Println(actual)
	}
}

func g(s string) string {
	return s + "fred"
}

func f(s string) string {
	return s + "simon"
}

