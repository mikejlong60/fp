package either

import (
	"errors"
	"reflect"
	"testing"
)

func TestStringIntTry(t *testing.T) {
	var actual = StringIntTry(stringToInt, "12")
	var expected = IntRight{12}
	if actual != expected {
		t.Errorf("StringIntTry  %v, want %v", actual, expected)
	}
}

func TestStringIntMapSuccess(t *testing.T) {
	var actual = StringIntMap(StringRight{"12"}, stringToInt)
	var expected = IntRight{12}
	if actual != expected {
		t.Errorf("StringIntMap had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func TestStringIntMapError(t *testing.T) {
	var actual = StringIntMap(StringRight{"12"}, stringToIntError)

	if reflect.TypeOf(actual).Name() != "Left" {
		t.Errorf("StringIntMap should have been type: Left  but was type:  %T  value: %v ", actual, actual)
	}
}

func TestStringIntFlatMapSuccess(t *testing.T) {
	var actual = StringIntFlatMap(StringRight{"12"}, stringToIntEither)
	var expected = IntRight{12}
	if actual != expected {
		t.Errorf("StringIntFlatMap had type: %T and value: %v, should have been type:  %T  value: %v ", actual, actual, expected, expected)
	}
}

func stringToInt(x string) (error, int) {
	return nil, 12
}

func stringToIntError(x string) (error, int) {
	return errors.New("bull"), -1
}

func stringToIntEither(x string) interface{} {
	return IntRight{12}
}
