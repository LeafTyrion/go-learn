package main

import (
	"fmt"
	"time"
)

//通常程序使用上锁的方式来保持资源的同步，避免资源竞争造成死锁
//go语言也支持上锁，但同时也有更好的解决方案
//使用管道 channel 则不需要进行上锁操作
//A往管道中写数据，B往管道中读数据，go会自动做好数据同步
func main() {

	//strChan:=make(chan string)//装字符串的管道
	numChan := make(chan int) //装数字的管道 无缓冲 写一个读一个
	//numChan := make(chan int,10)//装数字的缓冲管道 写一些读一些

	go func() {
		fmt.Println("读取数据")
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("子go程读取：", data)

		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println("子go程写入：", i)
			numChan <- i
		}
	}()

	for i := 20; i < 50; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("主go程写入：", i)
		numChan <- i
	}

	//子go程在管道为空的时候不会去读取数据
	//channel保证资源不被同一时刻访问，以免造成死锁
}
