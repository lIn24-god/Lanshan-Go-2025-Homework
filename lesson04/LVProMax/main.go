// LV PRO MAX
// 实现一个文件同步工具，编写一个文件同步工具，可以监视目录中的文件变化，并在文件修改时自动将变化同步到另一个目录。
//
// TIPS:
//
// 可以用 os.Stat() 获取文件信息
// 可以用 os.File.Readdir 获取目录下所有文件的信息
// os.FileInfo.Modtime 可以获取文件更新的时间，当然也可以通过别的方法判断文件是否被修改
// time.Ticker 实现周期事件，或者自己了解一下 cron

// ===================AI制造===========================

//// 伪代码框架
//func main() {
//    sourceDir := "要监视的目录"
//    targetDir := "同步到的目录"
//
//    // 创建定时器，比如每5秒检查一次
//    ticker := time.NewTicker(5 * time.Second)
//
//    for {
//        <-ticker.C // 等待下一个检查周期
//
//        // 1. 遍历源目录的所有文件
//        // 2. 对于每个文件，检查目标目录中对应的文件
//        // 3. 如果文件有变化（修改时间不同、大小不同或不存在），就复制
//        // 4. 可选：如果源目录删除了文件，目标目录也删除对应文件
//    }
//}

package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 检查文件是否需要同步
func needSync(sourcePath, targetPath string) bool {
	sourceInfo, err1 := os.Stat(sourcePath)
	if err1 != nil {
		// 源文件不存在，不需要同步
		return false
	}

	targetInfo, err2 := os.Stat(targetPath)
	if os.IsNotExist(err2) {
		// 目标文件不存在，需要同步
		return true
	}

	if err2 != nil {
		// 其他错误，需要同步
		return true
	}

	// 比较修改时间
	return !sourceInfo.ModTime().Equal(targetInfo.ModTime())
}

// 复制文件
func copyFile(sourcePath, targetPath string) error {
	// 打开源文件
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()

	// 创建目标文件
	target, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer target.Close()

	// 复制内容
	_, err = io.Copy(target, source)
	if err != nil {
		return err
	}

	return nil
}

// 获取目录中的所有文件
func getFiles(dirPath string) ([]os.FileInfo, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	// 读取目录中的所有文件
	files, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// 同步目录
func syncDirectory(sourceDir, targetDir string) {
	// 确保目标目录存在
	os.MkdirAll(targetDir, 0755)

	// 获取源目录文件列表
	sourceFiles, err := getFiles(sourceDir)
	if err != nil {
		fmt.Printf("读取源目录失败: %v\n", err)
		return
	}

	// 获取目标目录文件列表
	targetFiles, err := getFiles(targetDir)
	if err != nil && !os.IsNotExist(err) {
		fmt.Printf("读取目标目录失败: %v\n", err)
	}

	// 创建目标文件名的映射，用于快速查找
	targetMap := make(map[string]bool)
	for _, file := range targetFiles {
		targetMap[file.Name()] = true
	}

	// 同步源目录中的文件
	for _, file := range sourceFiles {
		sourcePath := sourceDir + "/" + file.Name()
		targetPath := targetDir + "/" + file.Name()

		if needSync(sourcePath, targetPath) {
			if file.IsDir() {
				// 如果是目录，递归同步
				os.MkdirAll(targetPath, 0755)
				syncDirectory(sourcePath, targetPath)
			} else {
				// 复制文件
				err := copyFile(sourcePath, targetPath)
				if err != nil {
					fmt.Printf("复制文件 %s 失败: %v\n", file.Name(), err)
				} else {
					fmt.Printf("同步文件: %s\n", file.Name())
				}
			}
		}
	}

	// 删除目标目录中不存在于源目录的文件
	sourceMap := make(map[string]bool)
	for _, file := range sourceFiles {
		sourceMap[file.Name()] = true
	}

	for _, file := range targetFiles {
		if !sourceMap[file.Name()] {
			targetPath := targetDir + "/" + file.Name()
			err := os.RemoveAll(targetPath)
			if err != nil {
				fmt.Printf("删除文件 %s 失败: %v\n", file.Name(), err)
			} else {
				fmt.Printf("删除文件: %s\n", file.Name())
			}
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("使用方法: %s <源目录> <目标目录>\n", os.Args[0])
		fmt.Printf("示例: %s ./source ./backup\n", os.Args[0])
		return
	}

	sourceDir := os.Args[1]
	targetDir := os.Args[2]

	// 验证源目录是否存在
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		fmt.Printf("错误: 源目录 '%s' 不存在\n", sourceDir)
		return
	}

	fmt.Printf("开始文件同步监视:\n")
	fmt.Printf("  源目录: %s\n", sourceDir)
	fmt.Printf("  目标目录: %s\n", targetDir)
	fmt.Printf("  检查间隔: 5秒\n")
	fmt.Printf("按 Ctrl+C 停止监视\n\n")

	// 立即执行一次初始同步
	syncDirectory(sourceDir, targetDir)
	fmt.Println("初始同步完成")

	// 创建定时器，每5秒检查一次
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// 开始监视循环
	for {
		select {
		case <-ticker.C:
			fmt.Printf("\n[%s] 检查文件变化...\n", time.Now().Format("15:04:05"))
			syncDirectory(sourceDir, targetDir)
		}
	}
}
