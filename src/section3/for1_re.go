//For문(1)

package main

import "fmt"

func main() {
	//반복문 - for
	//Go에서 유일하게 제공되는 반복문
	//다양한 사용법 숙지

	//예제1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	//에러발생1
	/*
	for i := 0; i < 5; i++
	{
		fmt.Println("ex1 : ", i)
	}
	*/
	
	//에러발생2
	/*
	for i := 0; i < 5; i++
		fmt.Println("ex1 : ", i)
	*/

	//예제2 (무한 루프)
	//초기화, 조건, 증가패턴을 다 제외하면 된다.
	/*
	for {
		fmt.Println("ex2 : Hello, Golang!")
		fmt.Println("ex2 : Infinite loop!")
	}
	*/

	//예제3 (range 용법)
	//range의 첫번째 값은 index, 두번째 값은 value
	loc := []string{"Seoul", "Busan", "Incheon", "Daegu"}
	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}