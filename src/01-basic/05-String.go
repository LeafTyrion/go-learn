package main

import "fmt"

//字符串
func main() {

	str := "HelloWorld"
	fmt.Printf("定义的字符串：%v \n", str)

	strings := `
1		Hello
1		World`
	fmt.Printf("原生的字符串：%v \n", strings)

	for i := 0; i < len(str); i++ {
		fmt.Printf("i: %v, v: %c \n", i, str[i])
	}

	const id = "123"
	fmt.Println("不能修改的字符串:", id)
}
