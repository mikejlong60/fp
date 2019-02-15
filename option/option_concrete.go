// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package option

import "fmt"

type IntSome struct {
	value int
}

type StringSome struct {
	value string
}

func StringIntTry(f func(string) (error, int), x string) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return IntSome{r}
	}
}

func StringIntMap(e interface{}, f func(string) (error, int)) interface{} {
	switch e.(type) {
	case None:
		return e
	case StringSome:
		v := e.(StringSome).value
		return StringIntTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or StringSome. Was type %T", e))
	}
}

func StringIntFlatMap(e interface{}, f func(string) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case StringSome:
		v := e.(StringSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or StringSome. Was type %T", e))
	}
}

func StringIntOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, IntSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or IntSome. Was type %T", other))
		}
	case StringSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type BoolSome struct {
	value bool
}

type StringSome struct {
	value string
}

func StringBoolTry(f func(string) (error, bool), x string) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return BoolSome{r}
	}
}

func StringBoolMap(e interface{}, f func(string) (error, bool)) interface{} {
	switch e.(type) {
	case None:
		return e
	case StringSome:
		v := e.(StringSome).value
		return StringBoolTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or StringSome. Was type %T", e))
	}
}

func StringBoolFlatMap(e interface{}, f func(string) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case StringSome:
		v := e.(StringSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or StringSome. Was type %T", e))
	}
}

func StringBoolOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, BoolSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or BoolSome. Was type %T", other))
		}
	case StringSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type FancyTypeSome struct {
	value FancyType
}

type StringSome struct {
	value string
}

func StringFancyTypeTry(f func(string) (error, FancyType), x string) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return FancyTypeSome{r}
	}
}

func StringFancyTypeMap(e interface{}, f func(string) (error, FancyType)) interface{} {
	switch e.(type) {
	case None:
		return e
	case StringSome:
		v := e.(StringSome).value
		return StringFancyTypeTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or StringSome. Was type %T", e))
	}
}

func StringFancyTypeFlatMap(e interface{}, f func(string) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case StringSome:
		v := e.(StringSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or StringSome. Was type %T", e))
	}
}

func StringFancyTypeOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, FancyTypeSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or FancyTypeSome. Was type %T", other))
		}
	case StringSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type IntSome struct {
	value int
}

type BoolSome struct {
	value bool
}

func BoolIntTry(f func(bool) (error, int), x bool) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return IntSome{r}
	}
}

func BoolIntMap(e interface{}, f func(bool) (error, int)) interface{} {
	switch e.(type) {
	case None:
		return e
	case BoolSome:
		v := e.(BoolSome).value
		return BoolIntTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or BoolSome. Was type %T", e))
	}
}

func BoolIntFlatMap(e interface{}, f func(bool) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case BoolSome:
		v := e.(BoolSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or BoolSome. Was type %T", e))
	}
}

func BoolIntOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, IntSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or IntSome. Was type %T", other))
		}
	case BoolSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type BoolSome struct {
	value bool
}

type BoolSome struct {
	value bool
}

func BoolBoolTry(f func(bool) (error, bool), x bool) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return BoolSome{r}
	}
}

func BoolBoolMap(e interface{}, f func(bool) (error, bool)) interface{} {
	switch e.(type) {
	case None:
		return e
	case BoolSome:
		v := e.(BoolSome).value
		return BoolBoolTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or BoolSome. Was type %T", e))
	}
}

func BoolBoolFlatMap(e interface{}, f func(bool) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case BoolSome:
		v := e.(BoolSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or BoolSome. Was type %T", e))
	}
}

func BoolBoolOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, BoolSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or BoolSome. Was type %T", other))
		}
	case BoolSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type FancyTypeSome struct {
	value FancyType
}

type BoolSome struct {
	value bool
}

func BoolFancyTypeTry(f func(bool) (error, FancyType), x bool) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return FancyTypeSome{r}
	}
}

func BoolFancyTypeMap(e interface{}, f func(bool) (error, FancyType)) interface{} {
	switch e.(type) {
	case None:
		return e
	case BoolSome:
		v := e.(BoolSome).value
		return BoolFancyTypeTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or BoolSome. Was type %T", e))
	}
}

func BoolFancyTypeFlatMap(e interface{}, f func(bool) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case BoolSome:
		v := e.(BoolSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or BoolSome. Was type %T", e))
	}
}

func BoolFancyTypeOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, FancyTypeSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or FancyTypeSome. Was type %T", other))
		}
	case BoolSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type IntSome struct {
	value int
}

type FancyTypeSome struct {
	value FancyType
}

func FancyTypeIntTry(f func(FancyType) (error, int), x FancyType) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return IntSome{r}
	}
}

func FancyTypeIntMap(e interface{}, f func(FancyType) (error, int)) interface{} {
	switch e.(type) {
	case None:
		return e
	case FancyTypeSome:
		v := e.(FancyTypeSome).value
		return FancyTypeIntTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or FancyTypeSome. Was type %T", e))
	}
}

func FancyTypeIntFlatMap(e interface{}, f func(FancyType) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case FancyTypeSome:
		v := e.(FancyTypeSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or FancyTypeSome. Was type %T", e))
	}
}

func FancyTypeIntOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, IntSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or IntSome. Was type %T", other))
		}
	case FancyTypeSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type BoolSome struct {
	value bool
}

type FancyTypeSome struct {
	value FancyType
}

func FancyTypeBoolTry(f func(FancyType) (error, bool), x FancyType) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return BoolSome{r}
	}
}

func FancyTypeBoolMap(e interface{}, f func(FancyType) (error, bool)) interface{} {
	switch e.(type) {
	case None:
		return e
	case FancyTypeSome:
		v := e.(FancyTypeSome).value
		return FancyTypeBoolTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or FancyTypeSome. Was type %T", e))
	}
}

func FancyTypeBoolFlatMap(e interface{}, f func(FancyType) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case FancyTypeSome:
		v := e.(FancyTypeSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or FancyTypeSome. Was type %T", e))
	}
}

func FancyTypeBoolOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, BoolSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or BoolSome. Was type %T", other))
		}
	case FancyTypeSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

type FancyTypeSome struct {
	value FancyType
}

type FancyTypeSome struct {
	value FancyType
}

func FancyTypeFancyTypeTry(f func(FancyType) (error, FancyType), x FancyType) interface{} {
	err, r := f(x)
	if err != nil {
		return None
	} else {
		return FancyTypeSome{r}
	}
}

func FancyTypeFancyTypeMap(e interface{}, f func(FancyType) (error, FancyType)) interface{} {
	switch e.(type) {
	case None:
		return e
	case FancyTypeSome:
		v := e.(FancyTypeSome).value
		return FancyTypeFancyTypeTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Interface type. Must be None or FancyTypeSome. Was type %T", e))
	}
}

func FancyTypeFancyTypeFlatMap(e interface{}, f func(FancyType) interface{}) interface{} {
	switch e.(type) {
	case None:
		return e
	case FancyTypeSome:
		v := e.(FancyTypeSome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Interface type.  Must be None or FancyTypeSome. Was type %T", e))
	}
}

func FancyTypeFancyTypeOrElse(e, other interface{}) interface{} {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, FancyTypeSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or FancyTypeSome. Was type %T", other))
		}
	case FancyTypeSome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}
