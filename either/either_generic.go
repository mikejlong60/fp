//go:generate genny -in=either_generic.go -out=either_concrete.go gen  "Either=interface{} AA=string,bool,FancyType BB=int"

package either

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type AA generic.Type
type BB generic.Type

type Either generic.Type

type BBRight struct {
	value BB
}

type AARight struct {
	value AA
}

func AABBTry(f func(AA) (error, BB), x AA) Either {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return BBRight{r}
	}
}

func AABBMap(e Either, f func(AA) (error, BB)) Either {
	switch e.(type) {
	case Left:
		return e
	case AARight:
		v := e.(AARight).value
		return AABBTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Either type. Must be Left or AARight. Was type %T", e))
	}
}

func AABBFlatMap(e Either, f func(AA) Either) Either {
	switch e.(type) {
	case Left:
		return e
	case AARight:
		v := e.(AARight).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Either type.  Must be Left or AARight. Was type %T", e))
	}
}

func AABBOrElse(e, other Either) Either {
	switch e.(type) {
	case Left:
		switch other.(type) {
		case Left, BBRight:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be Left or BBRight. Was type %T", other))
		}
	case AARight:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}
