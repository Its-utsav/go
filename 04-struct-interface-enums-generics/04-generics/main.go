package main

import "fmt"

func printSlice[T comparable, V string](slice []T, name V) {
	for _, value := range slice {
		fmt.Println(value, name)
	}
}

// type Stack[T int | string] struct {
// 	elements []T
// }

func main() {
	mySlice := []int{1, 2, 3, 4}
	languages := []string{"C", "C++", "Python", "Go", "JS/TS"}
	boolSlice := []bool{true, false, true, true}
	printSlice(mySlice, "Utsav")
	printSlice(languages, "Utsav")
	printSlice(boolSlice, "Utsav")
	// fmt.Println(languages)
	// stack := Stack[int]{
	// 	elements: []int{1, 2, 3},
	// }
	// fmt.Println(stack.elements)
}
