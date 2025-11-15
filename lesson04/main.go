/*package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("我的第一份作业.txt")
	if err != nil {
		fmt.Println("创建失败!", err)
		return
	}
	fmt.Fprintln(file, "姓名:林语杰")
	fmt.Fprintln(file, "2025-11-14")
	file.Close()
	fmt.Println("作业完成!")
}*/

//==========================================================

/*package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content, err := os.Open("我的第一份作业.txt")
	if err != nil {
		fmt.Println("打开文件失败!")
		return
	}
	defer content.Close()
	x, err := io.ReadAll(content)
	if err != nil {
		fmt.Println("读取文件失败!")
	}
	fmt.Println("文件内容：", string(x))
}*/

//====================================================================

/*package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content, err := os.Open("我的第一份作业.txt")
	if err != nil {
		fmt.Println("文件打开失败!")
		return
	}
	defer content.Close()
	file2, err := os.Create("作业备份.txt")
	if err != nil {
		fmt.Println("创建失败!")
		return
	}
	defer file2.Close()
	io.Copy(file2, content)
	fmt.Println("备份完成!")

}*/

//==================================================================

package main
