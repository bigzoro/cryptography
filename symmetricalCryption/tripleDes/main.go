package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func tripleDES(plainText []byte, key []byte) {
	// 获取3DES对象
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println("new triple des err: ", err)
		return
	}

	// 加密
	cipherText := make([]byte, len(plainText))
	block.Encrypt(cipherText, plainText)
	fmt.Printf("简单3DES加密后的数据：%s\n", cipherText)

	// 解密
	block.Decrypt(plainText, cipherText)
	fmt.Printf("简单3DES解密后的数据：%s\n", plainText)
}

func tripleDESCBCEncrypt(plainText []byte, key []byte) []byte {
	// 获取3DES对象
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println("new triple des err: ", err)
		return nil
	}

	// 对明文进行填充
	newPlainText := paddingLastGroup(plainText, block.BlockSize())
	// 获取CBC模式对象
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	cipherText := make([]byte, len(newPlainText))
	// 加密
	blockMode.CryptBlocks(cipherText, newPlainText)

	return cipherText
}

func tripleDESCBCDecrypt(cipherText []byte, key []byte) []byte {
	// 获取3DES对象
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println("new triple des err: ", err)
		return nil
	}

	// 获取CBC模式对象
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	// 解密
	blockMode.CryptBlocks(cipherText, cipherText)
	// 去除填充数据
	plainText := unPaddingLastGrooup(cipherText)

	return plainText
}

func tripleDESCTREncrypt(plainText []byte, key []byte) []byte {
	// 获取3DES对象
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println("new triple des err: ", err)
		return nil
	}

	// 创建CTR模式
	stream := cipher.NewCTR(block, key[:8])

	// 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

func tripleDESCTRDecrypt(cipherText []byte, key []byte) []byte {
	// 获取3DES对象
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		fmt.Println("new triple des err: ", err)
		return nil
	}

	// 创建CTR模式
	stream := cipher.NewCTR(block, key[:8])

	// 解密
	plainText := make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText
}

// 填充函数，输入明文、分组长度，输出填充后的数据
func paddingLastGroup(plainText []byte, blockSize int) []byte {
	// 求出最后一组剩余的字节数
	padNum := blockSize - len(plainText)%blockSize
	// 创建新的切片，长度等于padNum,
	char := []byte{byte(padNum)}
	// 切片创建闭并初始化
	newPlain := bytes.Repeat(char, padNum)
	// newPlain数组追加到原始明文的后面
	newText := append(plainText, newPlain...)
	return newText
}

// 去掉填充的数据
func unPaddingLastGrooup(plainText []byte) []byte {
	// 1. 拿去切片中的最后一个字节
	length := len(plainText)
	lastChar := plainText[length-1] //
	number := int(lastChar)         // 尾部填充的字节个数
	return plainText[:length-number]
}
func main() {

	key := []byte("123456781234567812345678")
	// 简单的3DES加解密，明文和密钥长度都要为168比特

	plainText1 := []byte("123456781234567812345678")
	tripleDES(plainText1, key)
	fmt.Println()

	// 使用CBC模式进行加密，可以加密任意长度的数据
	plainText2 := []byte("嘿哈")
	cipherText2 := tripleDESCBCEncrypt(plainText2, key)
	fmt.Printf("使用CBC模式加密后的数据：%s\n", cipherText2)
	plainText2 = tripleDESCBCDecrypt(cipherText2, key)
	fmt.Printf("使用CBC模式解密后的数据：%s\n", plainText2)
	fmt.Println()

	// 使用CTR模式进行加解密
	plainText3 := []byte("啊哈")
	cipherText3 := tripleDESCTREncrypt(plainText3, key)
	fmt.Printf("使用CTR模式加密后的数据：%s\n", cipherText3)
	plainText3 = tripleDESCTREncrypt(cipherText3, key)
	fmt.Printf("使用CTR模式解密后的数据：%s\n", plainText3)

}
