package main

import "fmt"

// Father 父类
type Father struct {
	name   string
	age    int
	gender string
}

func (father *Father) doSomething() {
	fmt.Printf("%v：打儿子", father.name)
}

// Combination 组合
type Combination struct {
	father Father //有字段名 则是组合
	score  float32
	name   string
}

// Son 子类
type Son struct {
	name   string
	Father //没有字段名 则是继承
	score  float32
}

// 方法重写
func (son *Son) doSomething() {
	fmt.Printf("%v玩游戏", son.name)
}

func main() {

	combination := Combination{
		father: Father{
			name:   "父亲",
			age:    50,
			gender: "男",
		},
		score: 100.0,
	}

	fmt.Println(combination)

	son := Son{}
	son.name = "儿子"
	son.age = 18
	son.gender = "男"
	son.score = 100.01

	fmt.Println(son)
	son.doSomething()
}
