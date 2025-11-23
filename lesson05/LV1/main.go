// Lv1
//
//	动手敲一下课件里面的代码，尝试理解一下（不用提交）
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	num := 0
	for range 10 {
		go func() {
			for range 1000000 {
				num++
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(num)

	//=======================================================

	sum := atomic.Int64{}
	for range 10 {
		go func() {
			for range 1000000 {
				sum.Add(1)
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(sum.Load())
}
