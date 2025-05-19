# Pointers
- pointer is a variable that store the address of the another variable  
```go
package main

import "fmt"

func changeNum(num int) {
	num = 100
	fmt.Println(num) // 100
}
func main() {
	var num int = 10
	changeNum(num)
	fmt.Println(num) // 10
}
```
- on above we try to change num value in `changeNum` function which is given from the main function 
- Why when we pass value in function than , received function received the copy of the value , not direct value


```go
package main

import "fmt"
func changeNum(num *int) {
	fmt.Println(*num, " will change to 69")
	*num = 69 // Dereferences for accessing the actual value
}

func main() {
	var num int = 10
	changeNum(&num)
	fmt.Println(num)
}
```


why used ?


- ✅ Efficient Memory Usage: Avoids copying large data structures.
- ✅ Pass by Reference: Functions can modify variables directly. 
- ✅ Dynamic Data Structures: Used in linked lists, trees, etc.