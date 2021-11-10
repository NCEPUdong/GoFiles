package main

import (
	"fmt"
	// "unicode/utf8"
	// unicode/utf8 模块的作用是用于计算utf-8下字符串字面值长度
)

type cyphertext string
type encrypted string

func main() {

	const str cyphertext = "L fdph, L vdz, L frqtxhuhg"
	const str1 cyphertext = "今天吃lunch了吗?"

	// len(str) 是golang的内置函数, 输出的是str的按字节计算的长度
	// 在unicode中有8位/16位/32位字符, 因此输出结果不总是等于字面值长度
	// 当然, 在字符串为ASCII的128个字符以内的字符, 长度都是8位, 没有计算字节产生的长度问题
	str.backward3()
	str1.cypherencode()
}

func (str cyphertext) cypherencode() {
	for _, c := range str {
		char := c
		fmt.Printf("%c", func() rune {
			if char >= 'a' && char <= 'z' {
				char += 13
				if char > 'z' {
					char -= 26
				}
			}
			return char
		}())
	}
	fmt.Print("\n")
}

func (str cyphertext) backward3() {
	for _, c := range str {
		char := c
		fmt.Printf("%c", func() rune {
			if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) {
				return char
			}
			char -= 3
			return char
		}())
	}
	fmt.Print("\n")
}
