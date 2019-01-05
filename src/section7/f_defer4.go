//함수 Defer(4)
package main

import (
	"fmt"
)

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b")) //중첩 함수 주의!, start 함수가 먼저 호출 후 "in a" 출력 후 end 호출
	fmt.Println("in a")
}

func main() {
	//예제1
	a()
}
