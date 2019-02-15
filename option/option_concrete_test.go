package option

import (
	"errors"
	"reflect"
	"testing"
)

func TestStringIntTry(t *testing.T) {
	var actual = StringIntTry(stringToInt, "12")
	var expected = IntSome{12}
	if actual != expected {
		t.Errorf("StringIntTry  %v, want %v", actual, expected)
	}
}

func TestStringIntMapSuccess(t *testing.T) {
	var actual = StringIntMap(StringSome{"12"}, stringToInt)
	var expected = IntSome{12}
	if actual != expected {
		t.Errorf("StringIntMap had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func TestStringIntMapError(t *testing.T) {
	var actual = StringIntMap(StringSome{"12"}, stringToIntError)
	if reflect.TypeOf(actual).Name() != "None" {
		t.Errorf("StringIntMap should have been type: None  but was type:  %T  value: %v ", actual, actual)
	}
}

func TestStringIntFlatMapSuccess(t *testing.T) {
	var actual = StringIntFlatMap(StringSome{"12"}, stringToIntEither)
	var expected = IntSome{12}
	if actual != expected {
		t.Errorf("StringIntFlatMap had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func TestStringIntOrElseRight(t *testing.T) {
	var actual = StringIntOrElse(StringSome{"12"}, IntSome{12})
	var expected = StringSome{"12"}
	if actual != expected {
		t.Errorf("StringIntOrElse had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func TestStringIntOrElseLeft(t *testing.T) {
	var actual = StringIntOrElse(None{}, IntSome{12})
	var expected = IntSome{12}
	if actual != expected {
		t.Errorf("StringIntOrElse had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func TestStringIntOrElseBadOtherType(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("StringIntOrElse did not fail on bad other type")
		}
	}()
	StringIntOrElse(None{}, FancyType{
		lastName:  "long",
		firstName: "mike",
		ssn:       "11111111",
		age:       632,
	})
}

func stringToInt(x string) (error, int) {
	return nil, 12
}

func stringToIntError(x string) (error, int) {
	return errors.New("bull"), -1
}

func stringToIntEither(x string) interface{} {
	return IntSome{12}
}
