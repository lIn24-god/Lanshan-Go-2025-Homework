package main

import "fmt"

func Fact(n int) int {
	p := 1
	for i := 1; i <= n; i++ {
		p *= i
	}
	return p
}
func Average(sum int, count int) float64 {
	var p float64
	p = float64(sum) / float64(count)
	return p
}

// -----------------------------------------------------------
func main() {
	fmt.Println("Hello林语杰")
	fmt.Println("hello" + "林语杰")
	fmt.Print("hello")
	fmt.Print("林语杰\n")
	name := "林语杰"
	fmt.Printf("Hello, %s!\n", name)

	//-----------------------------------------------------

	const pi = 3.14
	var area float64
	var r float64 = 5
	area = pi * r * r
	fmt.Println("area:", area)

	//-----------------------------------------------------

	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += i
		fmt.Println(sum)
	}

	//----------------------------------------------------

	var n int
	fmt.Println("Input:")
	fmt.Scan(&n)
	fmt.Println("n! = ", Fact(n))

	//----------------------------------------------------

	fmt.Println("输入整数，输入0停止：")
	var x int = 1
	var total, count int
	for x != 0 {
		fmt.Scan(&x)
		total = total + x
		count++
	}
	fmt.Println("total:", total, "count:", count)
	var aver float64 = Average(total, count)
	if aver >= 60 {
		fmt.Printf("平均成绩为%.2f , 合格！", aver)
	} else {
		fmt.Printf("平均成绩为%.2f , 不合格!", aver)
	}

}
