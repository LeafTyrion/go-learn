package main

import (
	"01-basic/15-Init/test"
	"fmt"
)

func main() {
	test.Test()
}

func init() {
	fmt.Println("this is main init")
}
