package main

import (
	"fmt"
	"strings"
)

type cyphertext string
type keyword string

func main() {
	var (
		str1 cyphertext = "Hello World!"
		key  keyword    = "GOLANG"
	)
	fmt.Println(str1.textEncode(key.keywordSer()))
}

func (str cyphertext) textEncode(key []int) string {
	var (
		keychar     rune
		encodedText string
		newStr      string
	)
	newStr = strings.ToUpper(string(str))
	for i, char := range newStr {
		if i < 6 {
			keychar = rune(key[i])
		} else {
			keychar = rune(key[i%6])
		}
		if char > 'Z' || char < 'A' {
			encodedText += string(char)
		} else {
			if char+keychar > 'Z' {
				char = char - 26 + keychar
			} else {
				char += keychar
			}
			encodedText += string(char)
		}
	}
	return encodedText
}

func (k keyword) keywordSer() []int {
	var newKeyword []int
	for _, char := range k {
		char -= 65
		newKeyword = append(newKeyword, int(char))
	}
	return newKeyword
}
