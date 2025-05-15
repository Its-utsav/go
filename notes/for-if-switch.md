## for
- for keyword used for looping 
- No while loop syntax but implement using for keyword
**While Loop**
```go
for <condition>{

}
```
**For loop**

```go
for init; condition;updation{
    // block of code
}
```
- `break` and `continue` keyword same as other language
- `range` - mainly used for iterating on array , string  or map it is same as range in python.

```go
for variable := range tillIndex{
    // block of code
}
```
- it loop till tillIndex but it is not included

## If
- no round bracket in conditions 
```go
age := 17
if age >= 18 {
	fmt.Println("Adult")
} else if age == 17 {
	fmt.Println("Wait for one year to be adult")
} else {
	fmt.Println("Minor")
}
```

---
- age only accessible inside the if-else block
```go
if age := 18; age >= 18 {
	fmt.Println("Age is 18 +")
} else {
	fmt.Println("Age less than 18")
}
```
- go does not have any ternary operator (?:) so you have to go with traditional way (May be in future)


- in switch case we do not have to write break statement  

1. Normal Switch Express Switch
```go
var day int = 1
var dayName string
switch day {
case 1:
	dayName = "Monday"
case 2:
	dayName = "Tuesday"
case 3:
	dayName = "Wednesday"
case 4:
	dayName = "Thursday"
case 5:
	dayName = "Friday"
case 6:
	dayName = "Saturday"
default:
	dayName = "Sunday"
}
fmt.Println("Day Name is", dayName)
```
- Switch with multiple Case

```go
switch time.Now().Weekday() {
case time.Saturday, time.Sunday:
	fmt.Println("Weekend")
default:
	fmt.Println("Weekday Work Hard")
}
```
Type Switch
```go
whichTypeOfData := func(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("Interger")
	case string:
		fmt.Println("String")
	case bool:
		fmt.Println("Boolean")
	default:
		fmt.Println("Value of", t)
	}
}
whichTypeOfData(1)
whichTypeOfData(true)
whichTypeOfData("utsav")
whichTypeOfData(3.14)
```
Switch case without expression (Mimic If else)

```go
var num int = 10
switch {
case num > 0:
	fmt.Println("Number is positive")
case num < 0:
	fmt.Println("Number is negative")
case num == 0:
	fmt.Println("Number is zero")
}

// or 
switch {
	case num >= 0:
		fmt.Println("Number is Bigger than 0")
	default:
		fmt.Println("Number is less than 0")
}
```