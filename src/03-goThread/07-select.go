package main

import (
	"fmt"
	"time"
)

func main() {

	//当程序中有多个channel协同工作时，需要在channel被触发时执行不同的操作
	//使用 select 监听多个管道
	//select 与 switch 很像，但是分支条件必须是管道

	channel1 := make(chan int)
	channel2 := make(chan int)

	go func() {
		for true {
			fmt.Println("监听中.......")
			select {
			case data := <-channel1:
				fmt.Println("监听到channel 1:", data)
			case data := <-channel2:
				fmt.Println("监听到channel 2:", data)
				//}
			}
		}
	}()

	//向channel1中写入数据
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("channel1 write: ", i)
			channel1 <- i
			time.Sleep(1 * time.Second / 2)

		}
	}()

	//向channel2中写入数据
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("channel2 write: ", i)
			channel2 <- i
			time.Sleep(1 * time.Second)

		}
	}()

	time.Sleep(15 * time.Second)
}
