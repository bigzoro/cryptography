package main

import (
	"fmt"
	"strings"
)

// 置换逆序加密
func reverseOrder(plaintText string) string {
	var cipherText strings.Builder
	// 把明文逆序，得到密文
	for i := len(plaintText) - 1; i > 0; i-- {
		cipherText.WriteString(string(plaintText[i]))
	}
	return cipherText.String()
}
func main() {
	// 我要告诉你一个秘密
	plaintText := "woyaogaosuniyigemimi"
	cipherText := reverseOrder(plaintText)
	fmt.Println(cipherText)
}
