package main

import (
	"fmt"
)

func matrixEncrypt(plainText string, key []byte) string {
	// 把明文要求排成矩阵，设矩阵为4列多列

	// 计算需要多少行
	row := len(plainText) / 4
	if len(plainText)%4 != 0 {
		row++
	}

	// 创建密文矩阵
	var plainTextMatrix [][4]string = make([][4]string, row)

	var a, b int
	// 循环倒明文的最后
	for i := 0; i < len(plainText); i++ {
		plainTextMatrix[a][b] = string(plainText[i])

		b++
		if b%4 == 0 && b != 0 {
			b = 0
			a++
		}
	}

	// 如果明文不够一行，就填充明文矩阵
	excr := len(plainText) % 4
	if excr != 0 {
		paddingNum := 4 - excr
		fmt.Println("padin", paddingNum)
		for i := 0; i < paddingNum; i++ {
			plainTextMatrix[row-1][3-i] = "@"
		}
	}

	// 生成密文矩
	var cipherText string
	// var test string
	var cipherTextMatrix [][4]string = make([][4]string, row)
	for i := 0; i < 4; i++ {
		for j := 0; j < len(plainTextMatrix); j++ {
			cipherTextMatrix[j][i] = plainTextMatrix[j][key[i]-1]
			// 这里可以直接把密文矩阵转换为密文
			// test += plainTextMatrix[j][key[i]-1]
		}

	}

	// 密文矩阵转化为密文
	for i := 0; i < 4; i++ {
		for j := 0; j <= len(cipherTextMatrix)-1; j++ {
			cipherText += cipherTextMatrix[j][i]
		}
	}

	// 返回密文
	return cipherText
}

func matrixDecrypt(cipherText string, key []byte) string {
	row := len(cipherText) / 4
	if len(cipherText)%4 != 0 {
		row++
	}
	var cipherTextMatrix [][4]string = make([][4]string, row)

	var a, b int
	// 根据密文得到密文矩阵
	for j := 0; j < len(cipherText); j++ {
		cipherTextMatrix[a][b] = string(cipherText[j])
		a++
		if a == row {
			a = 0
			b++
		}
	}

	// 把密文矩阵转换成明文矩阵
	var plainTextMatrix [][4]string = make([][4]string, row)
	for i := 0; i < 4; i++ {
		for j := 0; j < len(cipherTextMatrix); j++ {
			plainTextMatrix[j][key[i]-1] = cipherTextMatrix[j][i]
		}
	}

	// 明文矩阵转换成明文
	var plainText string
	for i := 0; i < row; i++ {
		for j := 0; j < 4; j++ {
			plainText += string(plainTextMatrix[i][j])
		}

	}
	// 返回明文
	return plainText
}
func main() {
	// 明文
	plainText := "woyaogaosuniyigemimi"
	// 密钥
	key := []byte{2, 3, 1, 4}

	fmt.Printf("明文：%s, 长度：%d\n", plainText, len(plainText))
	cipherText := matrixEncrypt(plainText, key)
	fmt.Printf("密文：%s, 长度：%d\n", cipherText, len(cipherText))
	plainText = matrixDecrypt(cipherText, key)
	fmt.Printf("解密后的明文：%s, 长度：%d\n", plainText, len(cipherText))
}
