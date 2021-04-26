package main

import (
	"fmt"
	"os"
)

//defer 延迟关键字，修饰语句或函数，确保这条语句或函数执行完毕出栈时执行
//一般用于资源清理，解锁、关闭文件
//在同一个函数多次调用时，执行顺序类似于栈机制
func main() {

	readFile("F:/myCode/go/go-learn/src/01-Variate.go")

}

func readFile(fileName string) {
	f1, err := os.Open(fileName)

	//匿名函数，关闭文件资源
	defer func(f1 *os.File) {
		fmt.Println("关闭文件资源")
		_ = f1.Close()
	}(f1) //尾部的括号指调用

	if err != nil {
		fmt.Printf("%v 打开失败！\n", fileName)
		println(err)
		return
	}

	content := make([]byte, 1024)
	length, err := f1.Read(content)
	if err != nil {
		fmt.Printf("%v 读取失败！\n", fileName)
		println(err)
		return
	}
	fmt.Printf("文件长度：%v \n", length)
	fmt.Printf("文件内容：%v \n", string(content))

}
