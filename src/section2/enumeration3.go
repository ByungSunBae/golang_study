//열거형3
package main

import "fmt"

func main() {
	// 밑줄은 스킵처리
	// iota는 규칙성을 갖는다.
	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	fmt.Println("G : ", GOLD)
	fmt.Println("P : ", PLATINUM)
}
