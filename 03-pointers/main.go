package main

import "fmt"

// func changeNum(num int) {
// 	num = 100
// 	fmt.Println(num) // 100
// }

func changeNum(num *int) {
	fmt.Println(*num, " will change to 69")
	*num = 69 // Dereferences for accessing the actual value
}

func main() {
	var num int = 10
	// changeNum(num)
	// fmt.Println(num) // 10
	// fmt.Println(&num) // Address of num variable
	changeNum(&num)
	fmt.Println(num)
}
