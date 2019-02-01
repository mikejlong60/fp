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

type RightC struct {
	value C
}

type E = error
type A = int
type B = int
type C = string
type BB interface{} //experimenting with this as a return type from MapFunction

type MapFunction func(A) B

type FlatMapFunction func(A) Either

func Map(e Either, f MapFunction) Either { //I think I could switch on the function type as well to deal with the other return types.
	switch e.(type) {
	case Left:
		return Left{e.(Left).value}
	case RightA:
		v := e.(RightA).value
		return RightB{f(v)}
	case RightB:
		v := e.(RightB).value
		return RightB{f(v)}
	default:
		panic(fmt.Sprintf("Unknown type %v", e))
	}
}

func FlatMap(e Either, f FlatMapFunction) Either {
	switch e.(type) {
	case Left:
		return Left{e.(Left).value}
	case RightA:
		v := e.(RightA).value
		return f(v)
	case RightB:
		v := e.(RightB).value
		return f(v)
	default:
		return e
	}
}

//func Try(f FlatMapFunction) Either {
//	r, err := f
//}
