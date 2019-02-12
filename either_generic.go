//go:generate genny -in=either_generic.go -out=either_concrete.go gen  "Either2=interface{} O=string,bool,FancyType P=int"

package fp

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type O generic.Type
type P generic.Type

type Either2 generic.Type

type PRight struct {
	value P
}

type ORight struct {
	value O
}

func OPTry(f func(O) (error, P), x O) Either2 {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return PRight{r}
	}
}

func OPMap(e Either2, f func(O) (error, P)) Either2 {
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

func OPFlatMap(e Either2, f func(O) Either2) Either2 {
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
