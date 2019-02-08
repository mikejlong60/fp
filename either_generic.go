//go:generate genny -in=either_generic.go -out=either_concrete.go gen  Eith=string,int

package fp

import (
	"github.com/cheekybits/genny/generic"
)

type Eith generic.Type

type EithGeneric struct{}

type EitherContainer interface{}

type EithRight struct {
	value Eith
}

func (e *EithGeneric) Try2(f func(Eith) (error, Eith), x Eith) EitherContainer {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return EithRight{r}
	}
}
