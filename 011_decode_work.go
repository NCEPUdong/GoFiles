package main

import (
	"fmt"
)

type cyphertext string
type keyword string

func main() {
	var (
		str1 cyphertext = "CSOITEUIWUIZNSROCNKFD"
		key  keyword    = "GOLANG"
	)
	fmt.Println(str1.textDecode(key.keywordSer()))
}

func (str cyphertext) textDecode(key []int) string {
	var (
		keychar     rune
		decodedText string
	)
	for i, char := range str {
		if i < 6 {
			keychar = rune(key[i])
		} else {
			keychar = rune(key[i%6])
		}

		if char-keychar < 'A' {
			char = char + 26 - keychar
		} else {
			char -= keychar
		}
		// if char > 'z' {
		// 	char -= 26
		// }
		decodedText += string(char)
	}
	return decodedText
}

func (k keyword) keywordSer() []int {
	var newKeyword []int
	for _, char := range k {
		char -= 65
		newKeyword = append(newKeyword, int(char))
	}
	return newKeyword
}
