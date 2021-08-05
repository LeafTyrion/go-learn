package main

import (
	"fmt"
	"go-learn/src/01-basic/15-Init/test"
)

func main() {
	test.Test()
}

func init() {
	fmt.Println("this is main init")
}
