package dao

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const dataFile = "./data/users.json"

var database map[string]string

// 简化的初始化：有文件就加载，没有就创建空map
func init() {
	database = make(map[string]string)
	data, err := os.ReadFile(dataFile)
	if err != nil {
		// 文件不存在或其他错误，就当全新开始
		return
	}
	// 简单解析：直接解析JSON到map
	json.Unmarshal(data, &database)
}

// 简化的保存：直接保存整个map
func save() {
	// 确保目录存在
	os.MkdirAll(filepath.Dir(dataFile), 0755)
	data, _ := json.MarshalIndent(database, "", "  ")
	os.WriteFile(dataFile, data, 0644)
}

// 业务函数

// AddUser 添加用户
func AddUser(username, password string) {
	database[username] = password
	save() // 改完就存
}

// FindUser 查找用户验证密码
func FindUser(username, password string) bool {
	return database[username] == password
}

// ChangeUserPassword 修改用户密码
func ChangeUserPassword(username, oldPwd, newPwd string) bool {
	if database[username] == oldPwd {
		database[username] = newPwd
		save() // 改完就存
		return true
	}
	return false
}
