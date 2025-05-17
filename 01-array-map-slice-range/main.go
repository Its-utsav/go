package main

func main() {

	// ---------------------------- Array Start here ----------------------------
	// var nums [5]float32
	// nums[0] = 1
	// nums[1] = 10
	// nums[2] = 100
	// nums[3] = 1000
	// nums[4] = 10000
	// fmt.Println(nums)
	// fmt.Println(len(nums))
	// fmt.Println(nums[2])
	// var array1 = [5]int{1, 2, 3, 4, 5}
	// var array2 = [...]int{1, 2, 3, 4, 5}
	// array3 := [3]int{1, 2, 3}
	//
	// 2-D Array
	//
	// nums := [2][2]int{{1, 2}, {3, 4}}
	// fmt.Println(nums)
	//
	// Linear Search
	// nums := [...]int{8, 3, 4, 6, 1}
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] == 1 {
	// 		fmt.Println("Luck Number HEHE at ", i)
	// 	}
	// }
	//
	// fmt.Println(nums)
	// ---------------------------- Array end here ----------------------------

	// ----------------------------- Slice Start Here ------------------------
	// Literl Way to decalre Slice
	// var names []string
	// fmt.Println(len(names))
	// Slice from array
	// var nums = [5]int{1, 2, 3, 4, 5}
	// slice := nums[1:4]
	// using make function
	// names := make([]string, 5, 10) // 5 length 10 capacity
	// names := make([]string, 5) // 5 length 5 capacity
	// names := make([]int, 2, 2) // 2 length 2 capacity
	// // names[3] = "Utsav" // Throw erro due the above capacity fixed to 2
	// fmt.Println(names)
	// names = append(names, 1) // add at 3
	// names = append(names, 2) // 4
	// names = append(names, 3) // 5
	// fmt.Println("cap", cap(names), "len", len(names), "names slice", names)
	// names := make([]int, 0, 2) // 0 length 2 capacity
	// fmt.Println(names)
	// names = append(names, 1) // add at 0
	// names = append(names, 2) // 1 ----- 2 capacity reached
	// names = append(names, 3) // 2 ------- 2 (init capacity) * 2 -> total 4 capacity 3 capacity
	// names = append(names, 4) // 2 --- 4 capacity
	// names = append(names, 4) // 2 --- 8 (4 * 2) capacity
	// fmt.Println("cap", cap(names), "len", len(names), "names slice", names)
	// num1 := make([]int, 0, 5)
	// num1 = append(num1, 10)
	// num2 := make([]int, len(num1))
	// fmt.Println(num1, num2)
	// copy(num2, num1)
	// fmt.Println(num1, num2)
	// var nums = []int{1, 2, 3, 5, 6, 7, 8, 9, 10}
	// fmt.Println(nums[:6])
	// s1 := []int{1, 2, 3}
	// s2 := []int{1, 2, 4}
	// isSliceSame := slices.Equal(s1, s2)
	// if isSliceSame {
	// 	fmt.Println("Slice same")
	// } else {
	// 	fmt.Println("Slice are not same")
	// }
	// -----------------------------------------------------------------------

	// ---------------------------- Map start here --------------------------------
	// m := make(map[string]string)
	// m["name1"] = "Go Lang"
	// m["name2"] = "Python"
	// m["name3"] = "C++"
	// m["name4"] = "C"
	// fmt.Println(m["Number"])
	// m := make(map[int]int)
	// m[1] = 1
	// m[2] = 2
	// delete(m, 2)
	// fmt.Println((m))
	// age := map[string]int{"utsav": 18, "sam": 26}
	// age["rahul"] = 24
	// // clear(age)
	// fmt.Println(age)
	// value, exists := age["utsav"]
	// if exists {
	// 	fmt.Println("Value in map", value)
	// } else {
	// 	fmt.Println("Key is not in map")
	// }
	// m1 := map[string]int{"utsav": 18, "sam": 26}
	// m2 := map[string]int{"utsav": 8, "sam": 26}
	// fmt.Println(maps.Equal(m1, m2))
	// ---------------------------- Map end Here  ---------------------------------

	// ---------------------------- range start Here  ------------------------------
	nums := make([]int, 0)
	nums = append(nums, 10)
	nums = append(nums, 12)
	nums = append(nums, 33)
	// for i := 0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// fmt.Println(nums, sum)

	// age := map[string]int{"utsav": 18, "rahul": 19, "sam": 22}

	// for key, value := range age {
	// 	fmt.Println(key, value)
	// }

	// for key := range age {
	// 	fmt.Println(key)
	// }

	// for b, unicode := range "utsav" {
	// 	fmt.Println(b, unicode)
	// }

	// for b, unicode := range "🔥utsav" { // Inetrsting
	// 	fmt.Println(b, unicode)
	// 	fmt.Println(b, string(unicode))
	// }
	// ---------------------------- range end Here  ---------------------------------
}
