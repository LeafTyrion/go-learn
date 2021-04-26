package main

import "fmt"

//只支持 i++
//不支持地址加减
//不支持三目运算
//if-else 只能接收 true、false
func main() {

	var student Student
	var pointer = &student

	if pointer == nil {
		fmt.Println("学生空指针", pointer)
	} else {
		fmt.Println("学生不空指针", pointer)
	}

}

type Student struct {
	name string
}
