package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// 简单的AES加解密应用
func aesEncrypt(plainText []byte, key []byte) []byte {
	// 获取AES加密对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("new aes cipher err: ", err)
		return nil
	}

	// 加密
	cipherText := make([]byte, 16)
	block.Encrypt(cipherText, plainText)

	return cipherText
}

func aesDecrypt(cipherText []byte, key []byte) []byte {
	// 获取AES加密对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("new aes cipher err: ", err)
		return nil
	}
	// 解密
	plainText := make([]byte, 16)
	block.Decrypt(plainText, cipherText)

	return plainText
}

func aesCTREncrypt(plainText []byte, key []byte) []byte {
	// 获取AES对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("new aes cipher err: ", err)
		return nil
	}

	// 获取CTR模式对象
	iv := []byte("1234567812345678")

	stream := cipher.NewCTR(block, iv)

	// 加密
	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

func aesCTRDecrypt(cipherText []byte, key []byte) []byte {

	// 获取AES对象
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("new aes cipher err: ", err)
		return nil
	}

	iv := []byte("1234567812345678")
	stream := cipher.NewCTR(block, iv)

	plainText := make([]byte, 1000)
	stream.XORKeyStream(plainText, cipherText)

	return plainText
}

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

	// 简单的AES加密，如果使用不对数据填充，一次只能加密128字节的数据
	plainText1 := []byte("1234567812345678")
	// aes的密钥长度为128字节
	key1 := []byte("1234567812345678")
	cipherText1 := aesEncrypt(plainText1, key1)
	fmt.Printf("简单加密后的数据： %s\n", cipherText1)
	plainText1 = aesDecrypt(cipherText1, key1)
	fmt.Printf("简单解密后的数据： %s\n", plainText1)

	// 使用CTR模式进行加密，当使用模式加密时，可以加密任意字节的数据
	plainText2 := []byte("绝了绝了绝了")
	key2 := []byte("1234567812345678")
	cipherText2 := aesCTREncrypt(plainText2, key2)
	fmt.Printf("使用CTR模式加密后的数据： %s\n", cipherText2)
	plainText2 = aesCTRDecrypt(cipherText2, key2)
	fmt.Printf("使用CTR模式解密后的数据： %s\n", plainText2)
}
