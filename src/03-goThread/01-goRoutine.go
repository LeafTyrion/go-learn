package main

import (
	"fmt"
	"time"
)

//go程（go语言的多线程，但是占用资源更少，创建更方便）
func main() {

	//启功子go程
	go display()

	count := 1

	//匿名方法启动go程
	go func() {
		count := 1
		for {
			fmt.Println("匿名方法线程: ", count)
			count++
			time.Sleep(1 * time.Second)
		}
	}()

	//主go程
	for {
		fmt.Println("主线程: ", count)
		count++
		time.Sleep(1 * time.Second)
	}

}

//子go程
func display() {
	count := 1

	//子go程
	for {
		fmt.Println("子线程: ", count)
		count++
		time.Sleep(1 * time.Second)
	}
}
