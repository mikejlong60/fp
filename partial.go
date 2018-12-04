package fp

type fn1 func(string) string
type fn2 func(string, string) string


func Partial1(f fn2, s1 string) fn1 {
	return func(s2 string) string {
		return f(s1, s2)
	}
}

/*
func(x, y int) int { return x + y }}

func AndThen(a fnString, b fnString) fnString {
	return func(s string) string {
		return b(a(s))
	}
}



  def curry[A, B, C](f: (A, B) => C): A => (B => C) = a => b => f(a, b)

  def partial1[A, B, C](a: A, f: (A, B) => C): B => C = (b: B) => f(a, b)

  def uncurry[A, B, C](f: A => B => C): (A, B) => C = (a: A, b: B) => f(a)(b)

*/

