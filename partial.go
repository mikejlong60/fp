package fp

import "fmt"

func Partial1(f func(string, string) string, s1 string) (func(string) string) {
	return func(s2 string) string {
		return f(s1, s2)
	}
}

//Takes a function that takes a string and returns a string and returns a function that takes two strings and applies the given function to it.
type Add func(string, string) string

func (operator Add) Curry(i string) func(string) string {
	return func(j string) string {
		return operator(i, j)
	}
}

func Uncurry(operator func(string) (func (string, string) string)) (func(string, string) string)  {
	return func(s1 string, s2 string) string {
		//var crap = operator(s1)
		//var crap2 = crap(s1, s2)
		fmt.Println(s1 + ":" + s2)
		return (operator(s1)(s1,s2))//crap2
	}

}

func curriedMapper(operator func(int) int) (func([]int) []int) {
	return func(m []int) []int {
		for i, n := range m {
			m[i] = operator(n)
		}
		return m
	}
}


/*
  def curry[A, B, C](f: (A, B) => C): A => (B => C) = a => b => f(a, b)

  def partial1[A, B, C](a: A, f: (A, B) => C): B => C = (b: B) => f(a, b)

  def uncurry[A, B, C](f: A => B => C): (A, B) => C = (a: A, b: B) => f(a)(b)

*/

