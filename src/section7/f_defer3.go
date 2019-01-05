//함수 Defer(3)
package main

import (
	"fmt"
)

func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 : ", i)
	} //선입후출
}

func main() {
	//예제1
	stack()
}
