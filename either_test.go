package fp

import (
	"errors"
	"reflect"
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

		var actual = Map(RightA{a}, add1000)
		var expected = RightB{a + 1000}

		if actual != expected {
			t.Errorf("Either Map  %q, want %q", actual, expected)
		}
	}
}

func TestMapLeftEither(t *testing.T) {
	var actual = Map(Left{errors.New("sad")}, add1000)
	var expected = Left{errors.New("sad")}

	if actual.(Left).value.Error() != expected.value.Error() {
		t.Errorf("Either Map with Left  %q, want %q", actual.(Left).value.Error(), expected.value.Error())
	}
}

func TestFlatMapNestedEither(t *testing.T) {
	var actual = Map(FlatMap(RightA{1}, makeF), add1000)
	var expected = RightB{2001}

	if actual != expected {
		t.Errorf("Either Map  %q, want %q", actual, expected)
	}
}

func TestFlatMapError(t *testing.T) {
	var actual = Map(FlatMap(RightA{-100}, makeF), add1000)
	if reflect.TypeOf(actual).Name() != "Left" {
		t.Errorf("Either Map  %q, want Left", actual)
	}
}

//TODO  - Write tests around Map2, Map3, ...

func add1000(i A) (error, B) {
	if i < 0 {
		return errors.New("can't have a number less than 0.  This is just for illustration;)"), -1
	} else {
		return nil, i + 1000
	}
}

func makeF(x A) Either {
	return Try(add1000, x)
}
