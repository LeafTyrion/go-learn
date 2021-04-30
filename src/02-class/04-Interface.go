package main

import (
	"fmt"
	"reflect"
)

//接口
func main() {

	var i, j, k interface{}

	array := []string{"1", "2"}
	i = array
	fmt.Println("i代表切片数组：", i)

	age := 20
	j = age
	fmt.Println("j代表数字：", j)

	str := "hello"
	k = str
	fmt.Println("k代表字符串", k)

	//只能在switch中使用，不然会报错
	//fmt.Println(str.(type))
	//判断类型
	value, ok := k.(int)
	if ok {
		fmt.Println("k的类型是int：", value)
	} else {
		fmt.Println("k的类型不是int", value)
	}

	//判断变量的数据类型
	fmt.Println("age的类型是：", reflect.TypeOf(array))

	arr := make([]interface{}, 3)
	arr[0] = 1
	arr[1] = "HELLO"
	arr[2] = true

	for _, value := range arr {
		switch v := value.(type) {

		case int:
			fmt.Println("int:", v)
		case string:
			fmt.Println("string:", v)
		case bool:
			fmt.Println("bool:", v)
		}
	}

}
