package option

func IntIntTry(f func(int, int) (error, int), x, y int) Option {
	err, r := f(x, y)
	if err != nil {
		return None{}
	} else {
		return IntSome{r}
	}
}

func Map2(a, b Option, f func(int, int) (error, int)) Option {

	switch a.(type) {
	case IntSome:
		switch b.(type) {
		case IntSome:
			return IntIntTry(f, a.(IntSome).value, b.(IntSome).value)
		default:
			return b
		}
	default:
		return a
	}
}
