package main

import (
	_ "01-basic/15-Init/test" //只调用init函数，不调用其他的函数
	"fmt"
)

func main() {
	fmt.Println("只调用init函数")
}
