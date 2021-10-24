package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	plainTextFile  = "./plaintText.txt"
	cipherTextFile = "./cipherText.txt"
	crackTextFile  = "./crackText.txt"
)

/*
	加密
	filename: 要加密的明文地址
	des: 存放加密后密文的地址
	key: 密钥
*/
func encrypt(plainTextFile string, dest string, key int) {

	// 从文件中读取明文
	plaintText, err := ioutil.ReadFile(plainTextFile)
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}
	// 得到密文切片，方便偏移
	ciphertextSlice := []uint8{}
	for _, v := range plaintText {
		// 把明文偏移key个单位
		intv := v + uint8(key)
		ciphertextSlice = append(ciphertextSlice, intv)
	}
	// 把密文写入文件
	err = ioutil.WriteFile(dest, ciphertextSlice, 0666)
	if err != nil {
		fmt.Println("write file err: ", err)
		return
	}

	fmt.Println("successfully encrypt")
}

/*
	解密
	cipherTextFile：密文路径
	plaintTextFile：解密后明文的存放路径
	key：明文
*/
func decrypt(cipherTextFile string, plainTextFile string, key int) {
	// 读取密文
	cipherText, err := ioutil.ReadFile(cipherTextFile)
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}

	plainTextSlice := []uint8{}
	// 根据密钥解密
	for _, v := range cipherText {
		intV := v - uint8(key)
		plainTextSlice = append(plainTextSlice, intV)
	}

	// 转换成明文字符串
	plaintText := string(plainTextSlice)

	// 写入本地
	err = ioutil.WriteFile(plainTextFile, []byte(plaintText), 0666)
	if err != nil {
		fmt.Println("write file err: ", err)
		return
	}

	fmt.Println("successfully decrypt")

}

/*
	破解
	cipherTextFile：加密文件
	crackTextFile：破解后文件的存储路径
*/
// 暴力破解
func cracking(cipherTextFile string, crackTextFile string) {
	// 读取密文
	cipherText, err := ioutil.ReadFile(cipherTextFile)
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}

	crackedTextSlice := []uint8{}
	var crackedText string
	file, err := os.OpenFile(crackTextFile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	// 暴力破解密文
	for i := 1; i <= 26; i++ {
		for _, v := range cipherText {
			intV := v - uint8(i)
			crackedTextSlice = append(crackedTextSlice, intV)
		}
		crackedText = string(crackedTextSlice)
		_, err = writer.WriteString(crackedText + "\n")
		if err != nil {
			fmt.Println("write file err: ", err)
			return
		}
		writer.Flush()
		crackedTextSlice = []uint8{}
	}

	fmt.Println("successfully cracke")
}
func main() {
	var key int
	var option int
	for {
		fmt.Println("1. 加密")
		fmt.Println("2. 解密")
		fmt.Println("3. 破解")
		fmt.Println("4. 退出程序")
		fmt.Println("请选择要进行的操作（1-4）：")
		fmt.Scanf("%d\n", &option)
		switch option {
		case 1:
			fmt.Println("请输入密钥：")
			fmt.Scanf("%d\n", &key)
			encrypt(plainTextFile, cipherTextFile, key)
		case 2:
			fmt.Println("请输入密钥：")
			fmt.Scanf("%d\n", &key)
			decrypt(cipherTextFile, plainTextFile, key)
		case 3:
			cracking(cipherTextFile, crackTextFile)
		case 4:
			os.Exit(0)
		}
	}

}
