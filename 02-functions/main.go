package main

import "fmt"

// func sayHi() {
// 	fmt.Println("Hello")
// }

// func add(a int, b int) int {
// 	return a + b
// }

// If paramters have same types
// func add(a, b int) int {
// 	return a + b
// }

// func languages() (string, string, string) {
// 	return "C", "C++", "Python"
// }

// func doSomething(fn func(num int) int) int {
// 	fmt.Println(fn(2))
// 	return 1
// }
func doSomething(int) func(a int) int {
	return func(a int) int {
		return 10
	}
}
func main() {
	// lang1, lang2, _ := languages()
	// myName := func() string {
	// 	return "utsav"
	// }
	// fmt.Println(myName())
	// fmt.Println(lang1, lang2)
	// fmt.Println(languages())

	fmt.Println(doSomething(1)(1))
}
