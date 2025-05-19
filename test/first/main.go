package main

import "fmt"

func sumAll(nums ...int) int {
	if len(nums) == 1 { // only one so no need of loop
		return nums[0]
	}
	var total int = 0
	for _, num := range nums {
		total += num
	}
	return total
}
func multiplier(factor int) func(int) int { // answer 4
	return func(arg int) int {
		ans := factor * arg
		return ans
	}
}
func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
func greetAll(greeting string, names ...string) {
	for _, name := range names {
		fmt.Println(greeting, name)
	}
}
func main() {
	sumAll(1, 2, 3, 4) //  use 1
	nums := []int{1, 2, 3, 4}
	sumAll(nums...) // use 2 // even answer of 2
	sumAll()
	//
	fn := multiplier(2) // answer 4
	fmt.Println(fn(3))  // answer 4

	// answer 5
	fn2 := counter()
	fn2() // 1
	fn2() // 2
	fn2() // 2
	// due to closure function are removed from callstack but refrecne is still maintain
	greetAll("Hi", "utsav", "hevin", "mummy")
}
