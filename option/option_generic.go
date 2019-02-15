//go:generate genny -in=option_generic.go -out=option_concrete.go gen  "Option=interface{} AA=string BB=int"

package option

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type AA generic.Type
type BB generic.Type

type Option generic.Type

type BBSome struct {
	value BB
}

type AASome struct {
	value AA
}

func AABBTry(f func(AA) (error, BB), x AA) Option {
	err, r := f(x)
	if err != nil {
		return None{}
	} else {
		return BBSome{r}
	}
}

func AABBMap(e Option, f func(AA) (error, BB)) Option {
	switch e.(type) {
	case None:
		return e
	case AASome:
		v := e.(AASome).value
		return AABBTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown Option type. Must be None or AASome. Was type %T", e))
	}
}

func AABBFlatMap(e Option, f func(AA) Option) Option {
	switch e.(type) {
	case None:
		return e
	case AASome:
		v := e.(AASome).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown Option type.  Must be None or AASome. Was type %T", e))
	}
}

func AABBOrElse(e, other Option) Option {
	switch e.(type) {
	case None:
		switch other.(type) {
		case None, BBSome:
			return other
		default:
			panic(fmt.Sprintf("Unknown other type.  Must be None or BBSome. Was type %T", other))
		}
	case AASome:
		return e
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}
