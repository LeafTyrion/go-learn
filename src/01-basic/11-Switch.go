package main

import (
	"fmt"
)

// switch
func main() {

	numArray := []int{1, 2, 3, 4}

	for key, value := range numArray {

		//case 默认自带break
		//若相应继续执行 则使用 fallthrough 关键字
		switch value {
		case 1:
			fmt.Printf("key: %v, value: %v \n", key, value)
		case 2:
			fmt.Printf("key: %v, value: %v \n", key, value)
			fallthrough
		case 3:
			fmt.Printf("key: %v, value: %v \n", key, value)
		default:
			fmt.Println("匹配失败")
		}

	}

}
