//go:generate genny -in=either_generic.go -out=either_concrete.go gen  "Either=interface{} O=string,bool,FancyType P=int"

package either

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type O generic.Type
type P generic.Type

type Either generic.Type

type PRight struct {
	value P
}

type ORight struct {
	value O
}

func OPTry(f func(O) (error, P), x O) Either {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return PRight{r}
	}
}

func OPMap(e Either, f func(O) (error, P)) Either {
	switch e.(type) {
	case Left:
		return e
	case ORight:
		v := e.(ORight).value
		return OPTry(f, v)
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

func OPFlatMap(e Either, f func(O) Either) Either {
	switch e.(type) {
	case Left:
		return e
	case ORight:
		v := e.(ORight).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}
