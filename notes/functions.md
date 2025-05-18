# Functions

- function rouble block of code
- go program starting point is function , which is automatically called 
- first class citizen in go lang
```go
// simple function
func sayHi() {
	fmt.Println("Hello")
}
```
----
```go
// simple function
func add(a int, b int) int {
	return a + b
}
func add(a, b int) int {
	return a + b
}
```
- returning multiple values
```go
// ⚠️ Type order matters
func languages() (string, string, string) {
	return "C", "C++", "Python"
}
lang1, lang2, lang3 := languages()
```
- usually we return two thing 
1. function logic
2. Error from function 

- first class citizen in go lang

```go
myName := func() string {
	return "utsav"
}
fmt.Println(myName())


func doSomething(fn func(num int) int) int {
	fmt.Println(fn(2))
	return 1
}
func main(){
    fn := func(num int) int {
		return 10
	}
	doSomething(fn)
}
```

- return function from a function

```go
func doSomething(int) func(a int) int {
	return func(a int) int {
		return 10
	}
}
fmt.Println(doSomething(1)(1)) // doSomething function -> call -> print 10 
```

## Variadic Functions
- a function that takes `n` number of arguments like `fmt.Println`

```go
func sumOfRandomNumbers(nums ...int) int { // we can even set any
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
```

## closure in go
- similar to `JavaScript` 
- function remember its variables
- in go we achieved through `anonymous functions`

```go
func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	count := counter()
	fmt.Println(count()) // 1
	fmt.Println(count()) // 2
	fmt.Println(count()) // 3
	fmt.Println(count()) // 4
	fmt.Println(count()) // 5
	fmt.Println(count()) // 6
}
```