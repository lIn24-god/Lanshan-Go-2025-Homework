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
	currentTime := time.Now()
	fmt.Println("开始时间：", currentTime.Format("15:04:05"))
	for range 10 {
		go func() {
			for range 1000000 {
				num++
			}
		}()
	}
	fmt.Println("结束时间：", time.Now().Format("15:04:05"))
	fmt.Println("经过了：", time.Since(currentTime))
	time.Sleep(2 * time.Second)
	fmt.Println(num)

	//=======================================================

	sum := atomic.Int64{}
	currentTime2 := time.Now()
	fmt.Println("开始时间：", currentTime2.Format("15:04:05"))
	for range 10 {
		go func() {
			for range 1000000 {
				sum.Add(1)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("结束时间：", time.Now().Format("15:04:05"))
	fmt.Println("经过了：", time.Since(currentTime2))
	time.Sleep(2 * time.Second)
	fmt.Println(sum.Load())
}

//=====================================================
