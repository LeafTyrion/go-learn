package main

import "fmt"

func main() {

	numChan := make(chan int, 10)

	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("写入：", i)
			numChan <- i
		}
		close(numChan) //关闭管道，应该放在写端执行
	}()

	//遍历管道
	//for range不知道管道已经写完了，会一直等待，导致deadlock
	//需要手动关闭管道
	for v := range numChan {
		fmt.Println("读取：", v)
	}

	//判断管道是否关闭
	for {
		v, ok := <-numChan
		if !ok {
			fmt.Println("管道关闭")
			break
		}
		fmt.Println("v", v)

	}

}
