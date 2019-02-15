package option

import (
	"testing"
)

func TestMap2Success(t *testing.T) {
	var e1 = IntSome{12}
	var e2 = IntSome{120}
	f := func(a, b int) (error, int) {
		return nil, a + b
	}
	actual := Map2(e1, e2, f)
	expected := IntSome{132}
	if actual != expected {
		t.Errorf("Map2 had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func TestMap2Failure1st(t *testing.T) {
	var e2 = IntSome{12}
	var e1 = None
	f := func(a, b int) (error, int) {
		return nil, a + b
	}
	actual := Map2(e1, e2, f)
	if actual != e1 {
		t.Errorf("Map2 had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, e1, e1)
	}
}

func TestMap2Failure2nd(t *testing.T) {
	var e1 = IntSome{12}
	var e2 = None
	f := func(a, b int) (error, int) {
		return nil, a + b
	}
	actual := Map2(e1, e2, f)
	if actual != e2 {
		t.Errorf("Map2 had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, e2, e2)
	}
}
