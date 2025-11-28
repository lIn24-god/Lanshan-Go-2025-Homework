// LV2andPLus
// 通过这节课的学习，相信你们都对 Go 语言的 I/O 有了比较清晰的认识。那么自己动手来实现一个日志 I/O 接口吧
//
// 需要实现的功能：
//
// 将日志信息保存在文件
// 日志信息中要包含时间和Unix时间戳

// LV2 PLUS
// 将日志信息同时输出文件和控制台中
//
// TIPS: io.MultiWriter
//
// 实现任何你想实现的功能：
//
// 通过环境变量，命令行参数（可以自己了解一下 flag 包），配置文件等方式指定日志文件路径
// 提供日志级别功能
// 允许自定义日志格式
// ……
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	//flag包,日志文件路径设定
	logPath := flag.String("log", "app.log", "日志文件路径")
	flag.Parse()
	// 打开一个日志文件，如果文件不存在则创建，追加写入
	file, err := os.OpenFile(*logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("OpenFile fail!")
		return
	}
	defer file.Close()

	// 创建multiWriter，同时输出到文件与控制台
	multiWriter := io.MultiWriter(file, os.Stdout)
	// 创建一个带时间戳的写入器
	logWriter := newTimestampWriter(multiWriter)

	// 模拟用户操作并记录日志
	fmt.Fprintln(logWriter, "用户登录")
	time.Sleep(2 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作A")
	time.Sleep(1 * time.Second)
	fmt.Fprintln(logWriter, "用户执行操作B")
}

// timestampWriter 是一个实现 io.Writer 接口的结构体，它在写入数据前添加时间和时间戳
type timestampWriter struct {
	logFile io.Writer
}

// 传入一个io.writer,file实现了io.writer接口
func newTimestampWriter(w io.Writer) *timestampWriter {
	return &timestampWriter{logFile: w}
}

func (tw *timestampWriter) Write(p []byte) (n int, err error) {
	// 添加时间戳和时间
	currentTime := time.Now()
	formatTime := currentTime.Format("2006-01-02 15:04:05")
	unixTime := currentTime.Unix()
	timestamp := fmt.Sprintf("[%s] [%d] %s", formatTime, unixTime, string(p))
	// 输出到文件
	return tw.logFile.Write([]byte(timestamp))
}
