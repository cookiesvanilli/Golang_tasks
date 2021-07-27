package main

import "fmt"

func main() {
	arrays()
	slices()
	slices2()
	appendSlice()
	copySlice()
	map1()
	map2()
	map3()
	map4()
	task1()
}
func arrays() {
	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
func slices() {
	x := make([]float64, 5, 10)
	fmt.Println(x)
}
func slices2() {
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:5]
	fmt.Println(x)
}
//append
func appendSlice() {
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}
//copy
func copySlice() {
	slice1 := []int{1,2,3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

// map - ассоциативный массив
func map1() {
	x:= make( map[string]int) //«x — это карта string-ов для int-ов»
	x["key"] = 10
	fmt.Println(x["key"])
}
func map2() {
	x:= make( map[int]int)
	x[1] = 10
	fmt.Println(x[1])
}
func map3() {
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
	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"]) // 0 или пустота
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}
func map4() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
		"Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
		"Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
		"B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
		"C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
		"N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
		"O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
		"F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
		"Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
func task1() {
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	min := x[0]
	for  _, a := range x {
		if (a < min) {
			min = a
		}
	}
	fmt.Println(min)
}