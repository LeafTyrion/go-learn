package main

import "fmt"

// 标签
func main() {

LABEL:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//重新进入循环，不会记录之前的状态
				//goto LABEL

				//跳到LABEL处继续循环，会记录之前的状态
				//continue LABEL

				//结束LABEL处的循环
				break LABEL

			}
			fmt.Println("i:", i, "j:", j)
		}
	}
}
