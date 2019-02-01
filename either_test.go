package fp

import (
	"errors"
	"strconv"
	"testing"
)

func TestMapRightEither(t *testing.T) {
	cases := []struct {
		actual int
	}{
		{1},
		{2},
	}
	for a, _ := range cases {

		var eit = RightA{a}
		var actual = Map(eit, add1000)
		var expected = RightB{a + 1000}

		if actual != expected {
			t.Errorf("Either Map  %q, want %q", actual, expected)
		}
	}
}

func TestMapLeftEither(t *testing.T) {
	var eit = Left{errors.New("sad")}
	var actual = Map(eit, add1000)
	var expected = Left{errors.New("sad")}

	if actual.(Left).value.Error() != expected.value.Error() {
		t.Errorf("Either Map with Left  %q, want %q", actual.(Left).value.Error(), expected.value.Error())
	}
}

func TestFlatMapNestedEither(t *testing.T) {
	var eit = RightA{1}
	var a1 = FlatMap(eit, boxedAdd1000)
	var actual = Map(a1, add1000)
	var expected = RightB{2001}

	if actual != expected {
		t.Errorf("Either Map  %q, want %q", actual, expected)
	}
}

func toBigString(i int) string {
	return strconv.Itoa(i + 1000)
}

func add1000(i int) int {
	return i + 1000
}

func boxedAdd1000(x int) Either {
	return RightB{x + 1000}
}
