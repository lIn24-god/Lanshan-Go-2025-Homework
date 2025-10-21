package main

import "fmt"

func main() {
	fmt.Println("Hello林语杰")
	fmt.Println("hello" + "林语杰")
	fmt.Print("hello")
	fmt.Print("林语杰\n")
	name := "林语杰"
	fmt.Printf("Hello, %s!\n", name)

	const pi = 3.14
	var area float64
	var r float64 = 5
	area = pi * r * r
	fmt.Println("area:", area)

	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += i
		fmt.Println(sum)
	}

}
