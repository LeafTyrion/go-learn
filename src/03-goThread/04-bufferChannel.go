package main

import (
	"fmt"
	"time"
)

func main() {

	//numchan:=make(chan int,10)//有缓冲的管道，容量为10
	//当缓冲区满，则写阻塞，只执行读取
	//当缓冲区空，则读阻塞，只执行写

	var names chan string
	names = make(chan string)

	go func() {
		fmt.Println("读取")
		data := <-names
		fmt.Println("names:", data)
	}()

	//只读不写打印空
	//只写不读会死锁 deadlock

	names <- "hello"

	time.Sleep(1 * time.Second)

	//阻塞在子go程时，会导致内存泄漏
	//阻塞在主go程时，会导致程序崩溃
}
