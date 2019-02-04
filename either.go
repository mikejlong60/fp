package fp

import "fmt"

type Either interface {
}

type Left struct {
	value E
}

type RightA struct {
	value A
}

type RightB struct {
	value B
}

type E = error
type A = int
type B = int
type C = string

type MapFunction func(A) B

type FlatMapFunction func(A) Either

func Map(e Either, f AToErrOrB) Either { //I think I could switch on the function type as well to deal with the other return types.
	switch e.(type) {
	case Left:
		return Left{e.(Left).value}
	case RightA:
		v := e.(RightA).value
		return Try(f, v)
	case RightB:
		v := e.(RightB).value
		return Try(f, v)
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

func FlatMap(e Either, f FlatMapFunction) Either {
	switch e.(type) {
	case Left:
		return Left{e.(Left).value}
	case RightA:
		v := e.(RightA).value
		return f(v)
		//	case RightB:
		//		v := e.(RightB).value
		//		return f(v)
	default:
		panic(fmt.Sprintf("Unknown type %T", e))
	}
}

//TODO implement Map2.  Here it is in Scala
//  def map2[EE >: E, B, C](b: Either[EE, B])(f: (A, B) => C) : Either[EE, C] = this.flatMap(a => b.map(b => f(a, b)))

//func Map2(a: Either, b: Either, f Map2F) Either {
//	FlatMap(a)
//}

type AToErrOrB func(A) (err error, r B)

func Try(f AToErrOrB, x A) Either {
	err, r := f(x)
	if err != nil {
		return Left{err}
	} else {
		return RightB{r}
	}
}
