package main

import "fmt"

//指针
func main() {

	name := "张三"
	ptr := &name
	fmt.Println("打印指针指向的值：", *ptr)
	fmt.Println("打印指针地址：", ptr)

	//定义指针类型
	namePtr := new(string)
	//给指针赋值
	*namePtr = "张三"
	fmt.Println("打印指针指向的值：", *namePtr)
	fmt.Println("打印指针地址：", namePtr)

	test := testPtr()
	fmt.Println("值：", *test)
	fmt.Println("指针：", test)

}

func testPtr() *string {
	city := "北京"
	ptr := &city
	return ptr
}
