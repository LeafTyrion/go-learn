package main

import "fmt"

//不定长数组，切片
func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	nums = append(nums, 10)
	fmt.Println("新增元素后的结果：", nums)

	slice := append(nums, 11)
	fmt.Println("新增元素后赋值给新的数组 slice：", slice)

	//长度指数组切片中的元素数量，容量指切片的总大小
	fmt.Printf("slice长度：%d，容量：%d \n", len(slice), cap(slice))

	//容量增长规则，一开始以2倍扩容，切片过大时不会以2倍增长，而是1.2
	var test []int
	for i := 0; i <= 50; i++ {
		fmt.Printf("测试切片的长度：%d,容量：%d \n", len(test), cap(test))
		test = append(test, i)
	}

	//切片操作，左闭右开，若切片越界以0补齐
	//切片结果为相当于原数组的引用，若修改切片数组里的内容，则原数组也会随之改变
	//若从0开始截取，则0可以省略;若截取到最后一个，则冒号右边可以省略；若全部，则都可省略
	newNums := nums[2:12]
	fmt.Println("切片后的新数组：", newNums)

	//若不想引用，则可以使用copy()函数，且目标数组也必须为切片
	var copyNums = []int{9, 8, 7}
	copy(copyNums, nums[:3])
	fmt.Println("拷贝后的数组：", copyNums)

	//直接截取字符串
	str := "helloWorld"[5:7]
	fmt.Println("截取后的字符串：", str)

	//使用make函数创建切片时指定长度和容量，无内容则空格替代（10个空格）
	//若没有指定容量，则默认和长度相等
	newSlice := make([]string, 10, 20)
	fmt.Printf("新切片的长度：%d，容量：%d \n", len(newSlice), cap(newSlice))

	//利用切片、make去copy数组
	cityArray := [3]string{"北京", "上海", "广州"}
	newCity := make([]string, len(cityArray))
	copy(newCity, cityArray[:])
	newCity = append(newCity, "深圳")
	fmt.Printf("原来的城市列表：%s，\n新增后的城市列表：%s \n", cityArray, newCity)

}
