package main

import (
	"fmt"
	"strings"
)

type (
	text      string
	wordList  []string
	frequency map[string]int
)

func main() {
	var (
		text1 text
	)
	text1 = `As far as eye could reach he saw nothing 
	but the stems of the great plants about him receding in the violet shade, 
	and far overhead the multiple transparency of huge leaves filtering 
	the sunshine to the solemn splendour of twilight in which he walked. 
	Whenever he felt able he ran again; the ground contined soft and springy, 
	covered with the same resilient weed 
	which was the first thing his hands had touched in Malacandra. 
	Once or twice a small red creature scuttled across his path, 
	but otherwise there seemed to be no life stirring in the wood; 
	nothing to fear-except the fact of wandering unprovisioned 
	and alone in a forest of unknown vegetation thousand 
	or millions of miles beyond the reach or knowledge of man.`

	for t, num := range wordCount(text1.preprocess()) {
		fmt.Printf("%-15s occurs %d times\n", t, num)
	}
}

func (t text) preprocess() wordList {
	var (
		list1 wordList
		list2 wordList
	)
	list1 = strings.Fields(strings.ToLower(string(t)))
	for _, word := range list1 {
		word = strings.Trim(word, " ")
		word = strings.Trim(word, ",")
		word = strings.Trim(word, ".")
		word = strings.Trim(word, ";")
		list2 = append(list2, word)
	}
	return list2
}

func wordCount(w wordList) frequency {
	// 除非使用了复合字面值来初始化map， 否则必须使用内置的make函数来为map分配空间
	// 创建map时， make函数接受一个或两个参数， 第二个参数用于预分配指定个数的key的空间， 默认map长度为0
	wordFrequency := make(frequency)
	for _, word := range w {
		wordFrequency[word]++
	}
	return wordFrequency
}
