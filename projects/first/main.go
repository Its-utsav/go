package main

import "fmt"

// Simple Inventory Tracker
// add into the inventory
// disply inventory
func addItem(inventory []string, item string) []string {
	inventory = append(inventory, item)
	return inventory
}

func displyInverntory(inventory []string) {
	fmt.Println("Printing Inventory")
	if len(inventory) == 0 {
		fmt.Println("Empty Inventory")
	}
	for _, item := range inventory {
		fmt.Println(item)
	}
}
func main() {
	inventory := make([]string, 0)

	fmt.Println("Welcome to the simple Inventory Tracker ")
	fmt.Println("Add products into the inventory and disply")

	// inventory = append(inventory, "Mouse")
	// inventory = append(inventory, "Keyboard")
	// inventory = append(inventory, "Laptop")
	inventory = addItem(inventory, "Mouse")
	inventory = addItem(inventory, "Keyboard")
	inventory = addItem(inventory, "Laptop")

	// for i := 0; i < len(inventory); i++ {
	// 	fmt.Println(inventory[i])
	// }
	displyInverntory(inventory)
}
