package main

import (
	"fmt"
	_ "go-learn/src/01-basic/15-Init/test" //只调用init函数，不调用其他的函数
)

func main() {
	fmt.Println("只调用init函数")
}
