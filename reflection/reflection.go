// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection
// TODO return later
package main

func walk(x interface{}, fn func(input string)) {
	fn("I still can't believe South Korea beat Germany 2-0 to put them last in their group")
}
