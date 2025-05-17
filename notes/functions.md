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