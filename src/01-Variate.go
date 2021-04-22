package main

import "fmt"

//变量
func main() {

	//先声明变量名，指定数据类型后赋值
	var name string
	name = "name"
	fmt.Println(name)

	//先声明变量名，不指定数据类型赋值
	var word = "word"
	fmt.Println(word)

	//不需要var关键字声明，不指定数据类型自动推导，
	hello := "hello"
	fmt.Println(hello)

}
