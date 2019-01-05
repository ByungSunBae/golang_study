//IF문(1)

package main

import "fmt"

func main() {
	//제어문(조건문)
	//IF 문 : 반드시 Boolean 검사 -> 1, 0 (사용불가 : 자동 형 변환 불가)
	//소괄호 사용 X

	var a int = 20
	b := 20

	//예제1
	if a >= 15 {
		fmt.Println("15이상")
	}

	//예제2
	if b >= 25 {
		fmt.Println("25이상")
	}

	//에러 발생1
	// 컴파일링시 바이트코드로 만들때 라인끝에 ;를 붙여버리기 때문에
	// if문 만들때 중괄호를 아래에 쓰면 안된다.
	/*
		if b >= 25
		{
			fmt.Println("25이상")
		}
	*/

	//에러 발생2
	/*
		if b >= 25
			fmt.Println("25이상")
	*/

	//에러 발생3
	/*
		if c := 1; c {
			fmt.Println("25이상")
		}
	*/
	/*
		에러X
		if c := true; c {
				fmt.Println("25이상")
			}
	*/

	// 예제3
	if c := 40; c >= 35 {
		fmt.Println("35이상")
	}

	// c += 20 // 에러발생
}
