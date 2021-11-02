package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	privateKeyFile = "./privateKey.pem"
	publicKeyFile  = "./publicKey.pem"
)

func GenerateKeyPair(bits int) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("generate private key err: ", err)
		return
	}

	// 通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)

	// 将私钥字符串设置到pem格式块中
	block := pem.Block{
		Type:    "PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}

	// 写入文件
	file1, err := os.Create(privateKeyFile)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	defer file1.Close()

	err = pem.Encode(file1, &block)
	if err != nil {
		fmt.Println("pem encode err: ", err)
		return
	}

	// 从得到的私钥对象中将公钥信息取出
	publicKey := privateKey.PublicKey
	// 通过x509标准将得到 的rsa公钥序列化为字符串
	publicKeyDer := x509.MarshalPKCS1PublicKey(&publicKey)
	// 将公钥字符串设置到pem格式块中
	block2 := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}
	// 写入文件
	file2, err := os.Create(publicKeyFile)
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	err = pem.Encode(file2, &block2)
	if err != nil {
		fmt.Println("write public key err: ", err)
		return
	}
}

// 公钥加密
func PublicKeyEncrypt(publicKeyFile string, plainText []byte) []byte {
	// 读取私钥
	publicKeyPem, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		fmt.Println("read public key file err: ", err)
		return []byte{}
	}

	// 解码成x509格式，rest是未解码完的数据存储在这里
	block, _ := pem.Decode(publicKeyPem)
	// 解析一个DER编码的公钥
	publicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		fmt.Println("x509 parse public key err: ", err)
		return nil
	}

	// 加密
	cipherData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		fmt.Println("public encrypt err: ", err)
		return nil
	}

	return cipherData
}

func PrivateKeyDecrypt(privateKeyFile string, cipherData []byte) []byte {
	// 读取私钥
	privateKeyPem, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		fmt.Println("read private file err: ", err)
		return nil
	}

	// 从数据中查找到PEM格式的块
	privateKeyDer, _ := pem.Decode(privateKeyPem)

	// 解析一个DER格式的私钥
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyDer.Bytes)
	if err != nil {
		fmt.Println("parse private key err: ", err)
		return nil
	}

	// 解密
	plaintData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherData)
	if err != nil {
		fmt.Println("private key decrypt err: ", err)
		return nil
	}

	return plaintData
}

func main() {
	bits := 1024

	// 生成公私钥
	GenerateKeyPair(bits)

	plainText := []byte("双击666")

	cipherData := PublicKeyEncrypt(publicKeyFile, plainText)
	fmt.Printf("加密后的数据: %s\n", cipherData)

	plainData := PrivateKeyDecrypt(privateKeyFile, cipherData)
	fmt.Printf("解密后的数据：%s\n", plainData)

}
