package main

import "fmt"

//定长数组，不可用动态改变长度
func main() {

	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("数组长度", len(nums))
	fmt.Println("数组内容", nums)

	//第一种方式变量数组
	for i := 0; i < len(nums); i++ {
		fmt.Printf("i: %d, num: %d \n", i, nums[i])
	}

	//获取数组下标和对应的值
	for key, value := range nums {
		if value != 4 {
			continue
		}
		fmt.Printf("key: %d, value: %d \n", key, value)
	}

	//忽略数组下标
	for _, v := range nums {
		fmt.Print(v)
	}
}
