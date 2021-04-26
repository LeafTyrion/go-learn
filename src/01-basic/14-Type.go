package main

import "fmt"

// Person 结构体定义
type Person struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {

	person := Person{
		name:   "张三",
		age:    18,
		gender: "男",
		score:  100,
	}

	fmt.Printf("张三的信息：%v \n", person)
	fmt.Printf("姓名：%v，年龄：%v，性别：%v，分数：%v \n", person.name, person.age, person.gender, person.score)

	pointer := &person
	fmt.Printf("使用指针 poniter 打印姓名：%v \n", pointer.name)
	fmt.Printf("使用指针 *poniter 打印姓名：%v \n", (*pointer).name)

	//全参赋值可以不写属性名,非全参赋值一定要带属性名
	a := Person{
		"张三",
		18,
		"女",
		99,
	}
	fmt.Printf("全参赋值可以不写属性名：%v\n", a)

}
