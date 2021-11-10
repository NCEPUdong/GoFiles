package main

import (
	"fmt"
)

func main() {
	var str1 string = "true"
	var str2 string = "no"
	var str3 string = "0"
	var str4 string = "go away"

	boolTrans := func(str string) interface{} {
		if str == "true" || str == "yes" || str == "1" {
			return true
		} else if str == "false" || str == "no" || str == "0" {
			return false
		} else {
			return "Error: type error!!"
		}
	}
	fmt.Println(boolTrans(str1))
	fmt.Println(boolTrans(str2))
	fmt.Println(boolTrans(str3))
	fmt.Println(boolTrans(str4))
}
