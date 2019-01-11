//함수 기초(4)
package main

import "fmt"

func multiply1(x, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func multiply2(x, y int) (int, int) {
	return x * 10, y * 20
}

// func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
// 	return a * 1, b * 2, c * 3, d * 4, e * 5
// }

func main() {
	//리턴 값 변수 사용
	//예제1
	a, b := multiply1(10, 5)
	fmt.Println("ex1 : ", a, b)

	//예제2
	c, d := multiply2(10, 5)
	fmt.Println("ex1 : ", c, d)
}
