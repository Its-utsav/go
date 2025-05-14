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