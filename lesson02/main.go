/*
import "fmt"

----------------

//复杂版  func function0(arr [10]int) map[int]int {
	result := make(map[int]int)
	for _, value := range arr {
		count := 0
		for _, value2 := range arr {
			if value == value2 {
				count++
			}
		}
		result[value] = count
	}
	return result
}
-------------
func function1(arr [10]int) map[int]int {
	result := make(map[int]int)
	for _, value := range arr {
		result[value]++
	}
	return result
}

func main() {
var arr [10]int
for index := range arr {
fmt.Scanln(&arr[index])
}
fmt.Println(function1(arr))
}*/

//_______________________________________________________________________________

package main

import "fmt"

type operator func(int, int) int

func CalculationOperator(op operator) func(int, int) int {
	return func(a int, b int) int {
		return op(a, b)
	}
}
func main() {
	var option string
	var op operator
	fmt.Scanln(&option)
	switch option {
	case "add":
		op = func(a, b int) int { return a + b }
	case "subtract":
		op = func(a, b int) int { return a - b }
	case "multiply":
		op = func(a, b int) int { return a * b }
	case "divide":
		op = func(a, b int) int {
			if b == 0 {
				fmt.Println("Input wrong!")
				return 0
			} else {
				return a / b
			}
		}
	}
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(CalculationOperator(op)(a, b))
}
