package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// 对少量数据进行hash
func Sha256Test1(data string) {
	// 计算散列值
	result := sha256.Sum256([]byte(data))
	fmt.Printf("少量数据hash后的数据：%x\n", result)
}

// 对文件进行hash
func Sha256Test2(fileName string) {
	// 获取文件句柄
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err：", err)
		return
	}
	defer f.Close()

	// 获取hash器
	hasher := sha256.New()
	_, err = io.Copy(hasher, f)
	if err != nil {
		fmt.Println("copy data err: ", err)
		return
	}

	// 计算散列值
	result := hasher.Sum(nil)

	fmt.Printf("文件数据hash后的数据：%x", result)
}

// 测试MD5
func main() {
	data := "123"
	filaName := "./text.txt"
	Sha256Test1(data)
	Sha256Test2(filaName)
}
