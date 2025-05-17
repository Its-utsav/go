package main

import "fmt"

func addItem(inventory map[string]int, product string, qty int) map[string]int {
	inventory[product] = qty
	return inventory
}

func display(inventory map[string]int) {
	fmt.Println("Printing Inventory")

	if len(inventory) == 0 {
		fmt.Println("Empty Inventory")
	}

	for product, qty := range inventory {
		fmt.Println(product, ":", qty)
	}
}

func getQty(inventory map[string]int, prodcut string) int {
	qty, isExists := inventory[prodcut]
	fmt.Println(qty, isExists)
	if isExists {
		return qty
	}
	return 0
}
func main() {
	inventory := make(map[string]int)
	addItem(inventory, "Laptop", 10)
	display(inventory)
	qty := getQty(inventory, "Mobile")
	if qty == 0 {
		fmt.Println("No Quantity found")
	}
}
