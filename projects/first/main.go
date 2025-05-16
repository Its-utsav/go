package main

import "fmt"

// Simple Inventory Tracker
// add into the inventory
// disply inventory

func main() {
	inventory := make([]string, 0)

	fmt.Println("Welcome to the simple Inventory Tracker ")
	fmt.Println("Add products into the inventory and disply")

	inventory = append(inventory, "Mouse")
	inventory = append(inventory, "Keyboard")
	inventory = append(inventory, "Laptop")

	for i := 0; i < len(inventory); i++ {
		fmt.Println(inventory[i])
	}
}
