//상수1
package main

import "fmt"

func main() {
	//상수
	//const 사용 초기화, 선언과 동시에 초기화해야함, 한 번 선언 후 값 변경 금지, 고정된 값 관리용
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	//const d = getHeight() 함수사용불가, 함수의 결과가 항상 같다는 보장이 없기 때문에!
	const e = 35.6
	const f = false
	/*
		에러발생
		const g string
		g = "Test3"
	*/

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)
}
