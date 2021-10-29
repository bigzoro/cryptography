package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func desCrypt(plainText []byte, key []byte) {
	// 创建DES加密对象
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 加密
	cipherText := make([]byte, len(plainText))
	block.Encrypt(cipherText, plainText)
	fmt.Printf("简单DES加密的数据：%s\n", cipherText)
	// 解密
	block.Decrypt(plainText, cipherText)
	fmt.Printf("简单DES加密的数据：%s\n", plainText)
}

func desCBCEncrypt(plainText []byte, key []byte) []byte {
	// 获取DES对象
	block, err := des.NewCipher(key)
	if err != nil {
		fmt.Println("new triple des err: ", err)
		return nil
	}

	// 对明文进行填充
	newPlainText := paddingLastGroup(plainText, block.BlockSize())
	// 获取CBC模式对象
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	// 加密
	cipherText := make([]byte, len(newPlainText))
	blockMode.CryptBlocks(cipherText, newPlainText)

	return cipherText
}

func desCBCDecrypt(cipherText []byte, key []byte) []byte {
	// 获取DES对象
	block, err := des.NewCipher(key)
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

func desCTREncrypt(plainText []byte, key []byte) []byte {
	// 获取DES对象
	block, err := des.NewCipher(key)
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

func desCTRDecrypt(cipherText []byte, key []byte) []byte {
	// 获取DES对象
	block, err := des.NewCipher(key)
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
	// 因为DES的密钥长度为64,所以一次只能加密8字节的数据
	key := []byte("12345678")

	// 简答加解密
	plainText1 := []byte("12345678")
	desCrypt(plainText1, key)

	// 使用CBC模式进行加解密
	plainText2 := []byte("12345678")
	cipherText2 := desCBCEncrypt(plainText2, key)
	fmt.Printf("使用CBC模式加密后的数据：%s\n", cipherText2)
	plainText2 = desCBCDecrypt(cipherText2, key)
	fmt.Printf("使用CBC模式解密后的数据：%s\n", plainText2)
	fmt.Println()

	// 使用CTR模式进行加解密
	plainText3 := []byte("啊哈")
	cipherText3 := desCTREncrypt(plainText3, key)
	fmt.Printf("使用CTR模式加密后的数据：%s\n", cipherText3)
	plainText3 = desCTREncrypt(cipherText3, key)
	fmt.Printf("使用CTR模式解密后的数据：%s\n", plainText3)

}
