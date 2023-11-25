package main

import "fmt"

func main() {

	//输出按照首字母顺序排序
	people := make(map[string]string)

	people["name"] = "zck"
	people["age"] = "18"
	people["sex"] = "男"
	people["live"] = "兰州"

	fmt.Println(people)

	fmt.Println(people["name"])

	var dog = map[string]string{
		"name":   "xzh",
		"age":    "180",
		"weight": "200",
		"height": "165",
	}

	for a, b := range dog {
		fmt.Println(a, b)
	}

}
