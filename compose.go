package fp

type fnString func(string) string

func Compose(a fnString, b fnString) fnString {
	return func(s string) string {
		return a(b(s))
	}
}

func AndThen(a fnString, b fnString) fnString {
	return func(s string) string {
		return b(a(s))
	}
}
