package main

import "fmt"

//主函数
func main() {

	v1, v2, v3 := test01(1, 1, "HelloWorld")
	fmt.Printf("test01函数的返回值:%v, %v, %v \n", v1, v2, v3)

	v1, v2, v3 = test02()
	fmt.Printf("test02函数的返回值:%v, %v, %v \n", v1, v2, v3)

	fmt.Printf("test03函数的返回值:%v \n", test03())

}

//函数定义，可以有多个返回值，返回值列表在参数列表之后
func test01(a int, b int, c string) (int, string, bool) {
	return a + b, c, a == b
}

//当返回值列表中有定义时，可以直接返回
func test02() (intRes int, strRes string, boolRes bool) {

	intRes = 100
	strRes = "HelloWorld"
	boolRes = false

	return
}

//返回值只有一个时则可以直接写类型并不加括号
func test03() string {
	return "test03"
}
