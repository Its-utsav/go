- [Array](#array)
- [Slice](#slice)
  - [Slice package](#slice-package)
  - [Length](#length)
  - [Capacity](#capacity)


## Array

- Array fixed side collection of same datatype , store element in continues memory location , access via index start from 0 , memory optimize
- If no values / elements were provided than array created with zero value
1. `int` -> 0
2. `bool` -> false
3. `string` -> empty string
4. `float64` -> 0 (Even for `float32`)

- Way to declare array
1. `var variableName[size] dataType`
2. `var variableName = [size]dataType{Elements}` -> allow partially to fully element initialization 
3. `var variableName = [...]dataType{Elements}` -> Infare the length of the array

- With shorthand
1. `variableName := [size]dataType{Elements}`
2. `variableName := [...]dataType{Elements}`
- `len(arrayVariable)` - return length of variable


2D array 

- `variableName := [row][col]{}`


## Slice
- Dynamic array in go lang
- Automatically grow or shrink when need
- Have serval built-in methods
- Most used construct in go lang
- Declaration is same as Array
- uninitialized slice are `nil`

```go
var names []string
fmt.Println(len(names))
```

- slice with `make` function
```go
nums := make([]string, 5) // != nil
```
1. `make([]int,length,capacity)`
2. `slice = append(slice,....newItems)`

- slice double the capacity when old capacity size limit reached
- slice copy
```go
num1 := make([]int, 0, 5) // length 0
num2 := make([]int, len(num1)) // num2 slice length 0
num1 = append(num1, 10) // append in num1 [10]
fmt.Println(num1, num2) // [10] []
copy(num2, num1) // copy num1 into num2 
fmt.Println(num1, num2) // [10] [0]
```
- why above `num2` slice length is 0 letter on `10` appending `num1` than copy `num1` into `num2`
- updated code
```go
num1 := make([]int, 0, 5)
num1 = append(num1, 10)
num2 := make([]int, len(num1))
fmt.Println(num1, num2)
copy(num2, num1)
fmt.Println(num1, num2)
```

- slicing in slice
```go
var nums = []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
fmt.Println(nums[1:]) // elements from 1 to till end
fmt.Println(nums[1:6]) // elements from 1 to 5 (second index not included)
fmt.Println(nums[:6]) // till 6 index but 6th index not included
```
### Slice package
- `slices.Equal(s1, s2)` - compare slice if length not same than return false , other wise check in increasing order 

### Length 
- return the number of element in slice or array
### Capacity
- return the total number of element slice can store before resize (It just double it)
- It is greater than `len`
- Used for memory Optimization
