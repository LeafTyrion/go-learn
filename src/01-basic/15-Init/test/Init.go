package test

import "fmt"

//test 函数没有返回值，没有参数，原型固定如下
//一个包中可以包含多个 test 函数，调用的顺序不确定
//init 不允许用户调用,通常用来初始化一些参数，如 MySQL 驱动，加载配置文件
//若执行使用 init，并不想使用其它函数，则可以使用_(下划线)
func init() {
	fmt.Println("this is first test function")
}

func init() {
	fmt.Println("this is second test function")
}

func Test() {
	fmt.Println("this is test function")
}
