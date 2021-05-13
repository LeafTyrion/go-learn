package main

import (
	"fmt"
	"time"
)

func main() {

	//单向管道，只读or只写
	//生产者消费者模式

	numChan := make(chan int, 50) //双向管道，可读可写

	go producer(numChan)
	go consumer(numChan)

	time.Sleep(2 * time.Second) //主go程睡一会儿要子go程执行
	//主协程直接结束，就不会有子协程输出了

}

//生产者，提供写 参数为一个只写的单向管道
func producer(writeChan chan<- int) {
	for i := 0; i < 50; i++ {
		fmt.Println("写入数据", i)
		writeChan <- i
	}
}

//消费者，提供读 参数为一个只读的单向管道
func consumer(readChan <-chan int) {

	for v := range readChan {
		fmt.Println("读取数据", v)
	}

	//for i := 0; i < 50; i++ {
	//	fmt.Println("读取数据", <-readChan)
	//}

}
