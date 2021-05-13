package main

import (
	"fmt"
	"os"
	"time"
)

// GOEXIT 提前退出当前go程
// return 返回当前函数（只返回 return 所在的那一层函数）
// exit 退出当前进程
func main() {

	go func() {
		func() {
			fmt.Println("内部函数")
			os.Exit(-1) //执行完内部函数直接退出当前main进程，不会执行 “结束”
			//runtime.Goexit() //只退出当前go程 执行 “结束”
		}()
		fmt.Println("外部函数")
	}()

	fmt.Println("主函数")
	time.Sleep(1 * time.Second)
	fmt.Println("结束")

}
