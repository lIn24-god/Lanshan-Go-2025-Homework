// LV PRO MAX
// 实现一个文件同步工具，编写一个文件同步工具，可以监视目录中的文件变化，并在文件修改时自动将变化同步到另一个目录。
//
// TIPS:
//
// 可以用 os.Stat() 获取文件信息
// 可以用 os.File.Readdir 获取目录下所有文件的信息
// os.FileInfo.Modtime 可以获取文件更新的时间，当然也可以通过别的方法判断文件是否被修改
// time.Ticker 实现周期事件，或者自己了解一下 cron
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	file, err := os.Stat("C:/GO--homework/lesson04/LV2andPlus/newnew")
	if err != nil {
		fmt.Println("stat fail")
		return
	}
	fmt.Println(file.Name())
	fmt.Println(file.IsDir())
	fmt.Println(file.Size())
	fmt.Println(file.ModTime().Format("2006-01-02 15:04:05"))
	time.Sleep(1 * time.Second)
	dirpath := filepath.Join("C:", "GO--homework", "lesson04", "LV2andPlus")
	dir, err := os.ReadDir(dirpath)
	if err != nil {
		fmt.Println("readDir fail")
		return
	}
	for _, entry := range dir {
		fmt.Println(" -", entry.Name())
	}
}
