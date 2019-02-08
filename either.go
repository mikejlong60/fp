package fp

import (
	"fmt"
)

type Either interface {
}

type Left struct {
	value E
}

type Right struct {
	value A
}

type E = error
type A = int

type MapFunction func(A) A

type FlatMapFunction func(A) Either

func Map(e Either, f AToErrOrA) Either { //I think I could switch on the function type as well to deal with the other return types.
	switch e.(type) {
	case Left:
		return Left{e.(Left).value}
	case Right:
		v := e.(Right).value
		return Try(f, v)
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

func FlatMap(e Either, f FlatMapFunction) Either {
	switch e.(type) {
	case Left:
		return Left{e.(Left).value}
	case Right:
		v := e.(Right).value
		return f(v)
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

//func ff func(a A) Either {/
//

//}
//func (a, b int, z float64) bool { return a*b < int(z) }

//TODO implement Map2.  Here it is in Scala
//func Map2(a Either, b Either, f func(A, A) A) Either {
//	var actual = FlatMap(a, func(b Either) Either {
//		return Map(b, func(b A) Either { return f(a, b) })
//	})
//
//	return actual
//}

//  def map2[EE >: E, B, C](b: Either[EE, B])(f: (A, B) => C) : Either[EE, C] = this.flatMap(a => b.map(b => f(a, b)))

//func Map2(a: Either, b: Either, f Map2F) Either {
//	FlatMap(a)
//}

type AToErrOrA func(A) (err error, r A)

func Try(f AToErrOrA, x A) Either {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return Right{r}
	}
}
