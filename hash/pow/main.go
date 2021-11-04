package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"time"
)

// 定义区块结构
type Block struct {
	//1.版本号
	Version uint64
	//2. 前区块哈希
	PrevHash []byte
	//3. Merkel根
	MerkelRoot []byte
	//4. 时间戳
	TimeStamp uint64
	//5. 难度值
	Bits uint64
	//6. 随机数，也就是挖矿要找的数据
	Nonce uint64
}

func Pow(block *Block) *Block {

	// 模拟目标值
	target := "0000100000000000000000000000000000000000000000000000000000000000"

	// 定义bigInt类型
	targetInt := big.Int{}
	// 将难度值转换成bigInt类型，指定为16进制格式
	targetInt.SetString(target, 16)

	// 计算出指定的哈希值
	var nonce uint64
	var hash [32]byte
	for {
		// 拼装区块数据，主要是nonce这个字段
		tempBlock := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Bits),
			// 这个值不断变换
			Uint64ToByte(nonce),
		}
		// 将二维切片转成一维切片
		blockInfo := bytes.Join(tempBlock, []byte{})

		// 做哈希运算
		hash = sha256.Sum256(blockInfo)

		// 将计算出的hash转成bigInt，好和目标值比较
		blockInt := big.Int{}
		blockInt.SetBytes(hash[:])

		// 与目标值比较，如果小于目标值，就代表找到了
		if blockInt.Cmp(&targetInt) == -1 {
			block.Nonce = nonce
			return block
		} else {
			// 没有找到，nonce值加1
			nonce++
		}

	}
}

//实现一个辅助函数，功能是将uint64转成[]byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer

	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buffer.Bytes()
}
func main() {

	// 模拟一个将要出块的区块
	block := &Block{
		Version:    0x1,
		PrevHash:   []byte("00000000000000000003de3ed5fae375845ae88a29448b5c0b5f4bc515276d10"),
		MerkelRoot: []byte("a329e36eefa678238489b24f5020f03e171e665d5e6df0f040918e31f668c7e4"),
		TimeStamp:  uint64(time.Now().Unix()),
		Bits:       21659344833264,
	}

	// 进行哈希运算（挖矿），并生成一个新的区块
	block = Pow(block)

	// 输出区块信息
	fmt.Printf("版本号：%v\n", block.Version)
	fmt.Printf("前区块哈希：%s\n", block.PrevHash)
	fmt.Printf("区块Merkel根：%s\n", block.MerkelRoot)
	fmt.Printf("时间戳：%v\n", block.TimeStamp)
	fmt.Printf("难度值：%v\n", block.Bits)
	fmt.Printf("随机数：%v\n", block.Nonce)
}
