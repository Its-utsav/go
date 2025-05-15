- [Array](#array)


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