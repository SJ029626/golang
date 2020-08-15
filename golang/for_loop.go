package main

import (
	"fmt"
)

func main() {
	var a int = 1
	for i := 10; i > 0; i-- {
		fmt.Println("value of a: ", a)
		a = a + 1
	}
}
