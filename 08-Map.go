package main

import "fmt"

//字典
func main() {

	//先定义，再初始化
	var aMap map[int]string
	aMap = make(map[int]string)
	fmt.Println(aMap)

	//map定义初始化,可以指定容量
	hashMap := make(map[int]string, 10)
	fmt.Println(hashMap)

	//map定义初始化并赋值
	stuMap := map[int]string{1: "张三", 2: "李四"}
	fmt.Println(stuMap)

	//遍历map
	for key, value := range stuMap {
		fmt.Printf("key: %d,value: %s \n", key, value)
	}

	//获取key对应的value
	//在map中不存在访问越界，所有key都有效
	//若key不存在，则返回零值：bool=false，数字=0，字符串=""
	stuName := stuMap[1]
	fmt.Println(stuName)

	//判断map中是否存在key
	key := 100
	value, isExist := stuMap[key]
	if isExist {
		fmt.Printf("%v,%d 存在，value：%s \n", isExist, key, value)
	} else if !isExist {
		fmt.Printf("%v,%d 不存在，value：%s \n", isExist, key, value)
	}

	//map插入
	stuMap[3] = "王五"
	fmt.Printf("插入后的map：%v \n", stuMap)

	//map删除,若key不存在，也不会报错
	delete(stuMap, 3)
	fmt.Printf("删除后的map：%v \n", stuMap)

}
