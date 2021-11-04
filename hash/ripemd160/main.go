package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/crypto/ripemd160"
)

// 对少量数据进行MD5
func Ripemd160Test1(data string) {
	// 获取hash器
	hahser := ripemd160.New()
	// 计算散列值
	result := hahser.Sum([]byte(data))
	fmt.Printf("少量数据hash后的数据：%x\n", result)
}

// 对大文件进行hash
func Ripemd160Test2(fileName string) {
	// 获取文件句柄
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err：", err)
		return
	}
	defer f.Close()

	// 获取hash器
	hasher := ripemd160.New()
	// 把文件内容拷贝到hash器
	_, err = io.Copy(hasher, f)
	if err != nil {
		fmt.Println("copy data err: ", err)
		return
	}

	// 计算散列值
	result := hasher.Sum(nil)

	fmt.Printf("大量数据hash后的数据：%x", result)
}

// 测试MD5
func main() {
	data := "123"
	filaName := "./text.txt"
	Ripemd160Test1(data)
	Ripemd160Test2(filaName)
}
