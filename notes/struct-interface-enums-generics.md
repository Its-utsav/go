- [Struct](#struct)
- [interface](#interface)
- [enums](#enums)
- [generics](#generics)

## Struct
- Struct is s user define data types that grouped different type of variable under the single name , in go their are no classes .

- struct declaration 
```go
type Order struct {
	id        int
	amount    float64
	status    string
	createdAt time.Time
}
```
```go
order1 := Order{
	id:        1,
	amount:    45.90,
	status:    "PENDING",
	createdAt: time.Now(),
}
var order2 Order = Order{1, 3.14, "pending", time.Now()}

order1.amount = 60.9 // modification by dot
fmt.Println(order1) // print entire order
```

Methods for struct

```go
func (order *Order) changeStatus(status string) {
	order.status = status // no need of Dereferences
}
func (order Order) getAmount() float64 {
	return order.amount 
}

// kind of constructor
func newOrder(id int, amount float64, status string) *Order { // Usually return via pointer
	myOrder := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}
```
---
```go
languages := struct {
	name   string
	isGood bool
}{"go lang", true}

```
## interface
## enums
## generics