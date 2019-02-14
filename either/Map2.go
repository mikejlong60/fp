package either

func IntIntTry(f func(int, int) (error, int), x, y int) Either {
	err, r := f(x, y)
	if err != nil {
		return Left{err}
	} else {
		return IntRight{r}
	}
}

func Map2(a, b Either, f func(int, int) (error, int)) Either {

	switch a.(type) {
	case IntRight:
		switch b.(type) {
		case IntRight:
			return IntIntTry(f, a.(IntRight).value, b.(IntRight).value)
		default:
			return b
		}
	default:
		return a
	}
}
