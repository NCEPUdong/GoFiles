package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		// 此处注意！ 想要字段能被导出为Json字段，必须要开头大写（即暴露出变量）
		// 通过为变量添加标签，可以修改Json导出时的字段名，XML字段名同理
		Name string  `json:"name"`
		Lat  float64 `json:"longitude"`
		Long float32 `json:"latitude"`
	}

	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	sample, _ := json.Marshal(locations[0])
	fmt.Println(string(sample))
	// {"name":"Bradbury Landing","longitude":-4.5895,"latitude":137.4417}
	for i := range locations {
		curiosity := locations[i]
		bytes, err := json.MarshalIndent(curiosity, "", "\t")
		/*
			MarshalIndent will format the Json
			{
			"name": "Bradbury Landing",
			"longitude": -4.5895,
			"latitude": 137.4417
			}
		*/
		exitOnError(err)
		fmt.Println(string(bytes))
	}
}

// exitOnError prints any errors and exits
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
