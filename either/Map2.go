package either

func Map2(a, b Either, f func(int, int) int) Either {

	switch a.(type) {
	case IntRight:
		switch b.(type) {
		case IntRight:
			return IntRight{f(a.(IntRight).value, b.(IntRight).value)}
		default:
			return b
		}
	default:
		return a
	}
}
