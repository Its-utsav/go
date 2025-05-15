package main

// Import package from standard Lib
// const name = "Go Lang"

// lol := 10 // Not work

// -------------- declare Multiple Constant ------ at same time
// const (
// 	port = 8000
// 	host = "localhost"
// )

func main() {
	// fmt.Println("Hello World!!!!") // Hello World

	// ---------------- Simple Values -------------------
	/*
		fmt.Println(1+1);
		fmt.Println(true,false)
		fmt.Println(3.14)
		fmt.Println(3.14 + 3.14)
	*/
	// --------------- Simple values end here -----------------

	// --------------- Variables start here -----------------
	/*
		// var name string = "Utsav"
		// ---- Simple Way
		// var name = "Utsav" // Infer type here
		// var isMale = true
		// var age int = 18

		// Shorthand syntax -> variable decalre and set some value
		// Used when we know the type of the value beforehand
		name := "Utsav"
		age := 18
		isMale := true
		macBookPrice := 999.9 // var macPrice fload64 := 999.9 //  var macPrice  := 999.9
		fmt.Println(name, age,isMale,macBookPrice)*/
	// --------------- Variables end here -----------------

	// --------------- Constants start here -----------------
	/*
		const name string = "Go Lang 2"
		// name = "Lol" // not allowed
		const age = 18
		fmt.Print(name)
	*/
	// --------------- Constants end here -----------------

	// -------------- for keyword start here --------------
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// -----------------
	// Infinte loop 1
	// for {
	// 	fmt.Print(true)
	// }
	// -----------------
	// Infinte loop 2
	// for true {
	// 	fmt.Print("Hello World")
	// }
	// -----------------
	// for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	// continue word
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	// break keyword
	// 	if i == 4 {
	// 		break
	// 	}
	// }
	// -------------- for keyword end here --------------

	// -------------- range keyword start here ------------
	// Loop till 99 ->  1-100
	// for i := range 100 {
	// 	fmt.Println(i)
	// }
	// -------------- range keyword end here --------------

	// ---------------- if-else start here ---------------
	// age := 17
	// if age >= 18 {
	// 	fmt.Println("Adult")
	// } else if age == 17 {
	// 	fmt.Println("Wait for one year to be adult")
	// } else {
	// 	fmt.Println("Minor")
	// }
	//
	// var user string = "ADMIN"
	// var hasAccess bool = true
	// if user == "ADMIN" || hasAccess {
	// 	fmt.Print("No permision")
	// } else {
	// 	fmt.Println("You have all permision")
	// }
	//
	// age only accessible inside the if-else block
	// if age := 18; age >= 18 {
	// 	fmt.Println("Age is 18 +")
	// } else {
	// 	fmt.Println("Age less than 18")
	// }
	// ---------------- if-else end here ---------------

	// ---------------- switch case start -------------
	//
	// Normal Switch Express Switch
	// var day int = 1
	// var dayName string
	// switch day {
	// case 1:
	// 	dayName = "Monday"
	// case 2:
	// 	dayName = "Tuesday"
	// case 3:
	// 	dayName = "Wednesday"
	// case 4:
	// 	dayName = "Thursday"
	// case 5:
	// 	dayName = "Friday"
	// case 6:
	// 	dayName = "Saturday"
	// default:
	// 	dayName = "Sunday"
	// }
	// fmt.Println("Day Name is", dayName)
	// ---
	//  Switch with multiple Case
	// ---
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Weekend")
	// default:
	// 	fmt.Println("Weekday Work Hard")
	// }
	// ----
	// Type Switch
	// ----
	// whichTypeOfData := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("Interger")
	// 	case string:
	// 		fmt.Println("String")
	// 	case bool:
	// 		fmt.Println("Boolean")
	// 	default:
	// 		fmt.Println("Value of", t)
	// 	}
	// }
	// whichTypeOfData(1)
	// whichTypeOfData(true)
	// whichTypeOfData("utsav")
	// whichTypeOfData(3.14)
	// ----
	// Switch case without expression (Mimic If else)
	// ----
	// var num int = 10
	// switch {
	// case num > 0:
	// 	fmt.Println("Number is positive")
	// case num < 0:
	// 	fmt.Println("Number is negative")
	// case num == 0:
	// 	fmt.Println("Number is zero")
	// }
	// ---------------- switch case end ---------------

	// ---------------- Array start here --------------------------

	// ---------------- Array end here --------------------------

	// name := "utsav"

	// fmt.Println("Hello world", name)
}
