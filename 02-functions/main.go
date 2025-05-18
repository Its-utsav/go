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
// func doSomething(int) func(a int) int {
// 	return func(a int) int {
// 		return 10
// 	}
// }

// func sumOfRandomNumbers(nums ...int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
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

	// fmt.Println(doSomething(1)(1))
	// nums := []int{1, 2, 3}
	// fmt.Println(sumOfRandomNumbers(nums...)) // kind of sperad operator in js
	count := counter()
	fmt.Println(count()) // 1
	fmt.Println(count()) // 2
	fmt.Println(count()) // 3
	fmt.Println(count()) // 4
	fmt.Println(count()) // 5
	fmt.Println(count()) // 6

}
