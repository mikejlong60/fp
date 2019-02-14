// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package either

import "fmt"

type IntRight struct {
	value int
}

type StringRight struct {
	value string
}

func StringIntTry(f func(string) (error, int), x string) interface{} {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return IntRight{r}
	}
}

func StringIntMap(e interface{}, f func(string) (error, int)) interface{} {
	switch e.(type) {
	case Left:
		return e
	case StringRight:
		v := e.(StringRight).value
		return StringIntTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be Left or StringRight. Was type %T", e))
	}
}

func StringIntFlatMap(e interface{}, f func(string) interface{}) interface{} {
	switch e.(type) {
	case Left:
		return e
	case StringRight:
		v := e.(StringRight).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be Left or StringRight. Was type %T", e))
	}
}

func StringIntOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case Left:
		switch other.(type) {
		case Left, IntRight:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be Left or IntRight. Was type %T", other))
		}
	case StringRight:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type BoolRight struct {
	value bool
}

func BoolIntTry(f func(bool) (error, int), x bool) interface{} {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return IntRight{r}
	}
}

func BoolIntMap(e interface{}, f func(bool) (error, int)) interface{} {
	switch e.(type) {
	case Left:
		return e
	case BoolRight:
		v := e.(BoolRight).value
		return BoolIntTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be Left or BoolRight. Was type %T", e))
	}
}

func BoolIntFlatMap(e interface{}, f func(bool) interface{}) interface{} {
	switch e.(type) {
	case Left:
		return e
	case BoolRight:
		v := e.(BoolRight).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be Left or BoolRight. Was type %T", e))
	}
}

func BoolIntOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case Left:
		switch other.(type) {
		case Left, IntRight:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be Left or IntRight. Was type %T", other))
		}
	case BoolRight:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type FancyTypeRight struct {
	value FancyType
}

func FancyTypeIntTry(f func(FancyType) (error, int), x FancyType) interface{} {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return IntRight{r}
	}
}

func FancyTypeIntMap(e interface{}, f func(FancyType) (error, int)) interface{} {
	switch e.(type) {
	case Left:
		return e
	case FancyTypeRight:
		v := e.(FancyTypeRight).value
		return FancyTypeIntTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be Left or FancyTypeRight. Was type %T", e))
	}
}

func FancyTypeIntFlatMap(e interface{}, f func(FancyType) interface{}) interface{} {
	switch e.(type) {
	case Left:
		return e
	case FancyTypeRight:
		v := e.(FancyTypeRight).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be Left or FancyTypeRight. Was type %T", e))
	}
}

func FancyTypeIntOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case Left:
		switch other.(type) {
		case Left, IntRight:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be Left or IntRight. Was type %T", other))
		}
	case FancyTypeRight:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}
