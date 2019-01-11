//상수1
package main

import "fmt"

func main() {
	//상수
	//const 사용 초기화, 선언과 동시에 초기화해야함, 한 번 선언 후 값 변경 금지, 고정된 값 관리용
	const a string = "String1"
	const b = "String2"
	const c int32 = 9 * 9
	// const d = getWidth() 함수 사용 불가, 함수의 결과가 항상 같다는 보장이 없기 때문이다.
	const e = 35.6
	const f = false
	/*
	에러발생
	const g string
	g = "String3"
	*/

	fmt.Println(
		"a : ", a,
		"b : ", b,
		"c : ", c,
		"e : ", e,
		"f : ", f,
	)
}