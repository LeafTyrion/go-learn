package main

import "fmt"

// Person 结构体模拟类，类中不能直接写方法
type Person struct {
	name string
	age  int
}

// Eat1 方法绑定到 Person 类,入参使用指针，相当于改变本身
func (p *Person) Eat1() {
	fmt.Println(p.name + "人吃饭")
	p.name = "指针"
}

// Eat2 方法绑定到 Person 类，入参使用类，相当于传入副本
func (p Person) Eat2() {
	fmt.Println(p.name + "吃饭")
	p.name = "非指针"
}

func main() {

	person := Person{
		"张三",
		18,
	}
	//方法绑定
	//person.Eat1()

	person1 := person
	person2 := person

	fmt.Println("使用非指针修改名字")
	person2.Eat2()
	fmt.Println(person2)

	fmt.Println("使用指针修改名字")
	person1.Eat1()
	fmt.Println(person1)

}
