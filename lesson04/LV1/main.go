//LV1
//借助这节课学习的 time 包和 bufio 包自己验证一下带缓冲 I/O 写是否比不带缓冲 I/O 写的效率更高（通常写入次数越多差异越明显）
//TIPS:用 time.Since记录花费的时间

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func NoBufferTest() {
	file, err := os.OpenFile("NoBuffer.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("OpenFile fail!")
		return
	}
	defer file.Close()
	CurrentTime := time.Now()
	fmt.Println("不带缓冲开始运行", CurrentTime.Format("15:04:05"))
	for range 1000000 {
		file.WriteString("loop")
	}
	fmt.Println("运行时长:", time.Since(CurrentTime), "当前时间:", time.Now().Format("15:04:05"))
}

func BufferTest() {
	file, err := os.OpenFile("Buffer.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("OpenFile fail!")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	CurrentTime := time.Now()
	fmt.Println("带缓冲开始运行", time.Now().Format("15:04:05"))
	for range 1000000 {
		writer.WriteString("loop")
	}
	writer.Flush()
	fmt.Println("运行时长:", time.Since(CurrentTime), "当前时间:", time.Now().Format("15:04:05"))
}

func main() {
	fmt.Println("test start!")
	NoBufferTest()
	time.Sleep(2 * time.Second)
	BufferTest()
}
