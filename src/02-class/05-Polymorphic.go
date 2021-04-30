package main

import "fmt"

//多态

type IAttack interface {
	Attack()
}

type Human01 struct {
	name  string
	level int
}

type human02 struct {
	name  string
	level int
}

func (human *human02) Attack() {
	fmt.Println("i am ", human.name, ", level is", human.level)
}

func (human *Human01) Attack() {
	fmt.Println("i am ", human.name, ", level is", human.level)
}

func main() {

	var interHuman IAttack

	human01 := Human01{
		"number01",
		1,
	}
	human01.Attack()
	//接口需要指针赋值
	interHuman = &human01
	interHuman.Attack()

	human02 := human02{
		"number02",
		2,
	}
	human02.Attack()
	interHuman = &human02
	interHuman.Attack()

	//同一个interfaceHuman可以接收不同的human类型
	DoAttack(&human01)
	DoAttack(&human02)

}

// DoAttack 定义一个函数,多态的通用接口，传入不同的类，表现出不同的形式
//本质都实现了Attack()方法
func DoAttack(a IAttack) {
	a.Attack()

}
