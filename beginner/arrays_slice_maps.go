package main

import "fmt"

func main() {
	example_array()
	example_slice()
	example_map()
}

func example_array() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	var total float64 = 0

	asz := len(x)
	for i := 0; i < asz; i++ {
		total += x[i]
	}
	fmt.Println("Average of array values: ", (total / (float64)(asz)))
}

func example_slice() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice2)
}

func example_map() {
	//Using make function create map object
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements)
	fmt.Println(elements["Li"])
	//using delete function we can delete keys and values
	//fmt.Println("Delete he key from map: ", delete(elements,"he"))

	//Other way to crate map
	elements1 := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B": "Boron",
		"C": "Carbon",
		"N": "Nitrogen",
		"O": "Oxygen",
		"F": "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements1)
	fmt.Println(elements1["H"])
}
