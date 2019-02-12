//go:generate genny -in=either_generic.go -out=either_concrete.go gen  "O=string,bool P=int"

package fp

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

type O generic.Type
type P generic.Type

type PRight struct {
	value P
}

type ORight struct {
	value O
}

func OPTry(f func(O) (error, P), x O) interface{} {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return PRight{r}
	}
}

func OPMap(e interface{}, f func(O) (error, P)) interface{} {
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

func OPFlatMap(e interface{}, f func(O) interface{}) interface{} {
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
