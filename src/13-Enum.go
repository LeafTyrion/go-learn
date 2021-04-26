package main

import "fmt"

//go中没有枚举，可以使用const+iota（常量累加器）来模拟
func main() {

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(WEDNESDAY)
	fmt.Println(THURSDAY)
	fmt.Println(FRIDAY)
	fmt.Println(SATURDAY)
	fmt.Println(SUNDAY)

}

//const 属于预编译期赋值，不需要:=进行自动推导类型
const (
	// MONDAY iota从0开始，每一行递增1
	MONDAY = iota + 1
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)
