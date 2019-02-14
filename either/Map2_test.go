package either

import (
	"errors"
	"testing"
)

func TestMap2Success(t *testing.T) {
	var e1 = IntRight{12}
	var e2 = IntRight{120}
	f := func(a, b int) int {
		return a + b
	}
	actual := Map2(e1, e2, f)
	expected := IntRight{132}
	if actual != expected {
		t.Errorf("StringIntFlatMap had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func TestMap2Failure1st(t *testing.T) {
	var e2 = IntRight{12}
	var e1 = Left{errors.New("Dang!")}
	f := func(a, b int) int {
		return a + b
	}
	actual := Map2(e1, e2, f)
	if actual != e1 {
		t.Errorf("StringIntFlatMap had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, e1, e1)
	}
}

func TestMap2Failure2nd(t *testing.T) {
	var e1 = IntRight{12}
	var e2 = Left{errors.New("Dang!")}
	f := func(a, b int) int {
		return a + b
	}
	actual := Map2(e1, e2, f)
	if actual != e2 {
		t.Errorf("StringIntFlatMap had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, e2, e2)
	}
}
