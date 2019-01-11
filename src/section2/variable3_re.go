//변수3
package main

import "fmt"

func main() {
	//짧은 선언 (Short Assignment Statement)
	//함수 안에서만 사용(전역X), 선언 후 할당 예외 발생
	//주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일수 있다 --> 추천
	// => 즉, := 으로 되어있으면 "함수내에서만 사용하는구나"라고 이해하면 된다.

	shortVar1 := 3
	shortVar2 := "String"
	shortVar3 := false

	//shortVar3 := true
	//짧은 선언 후 재선언 불가능... 예외발생!

	fmt.Println(
		"shortVar1 : ", shortVar1,
		"shortVar2 : ", shortVar2,
		"shortVar3 : ", shortVar3,
	)

	//Example
	if i := 10; i < 11 {
		fmt.Println("Short Variable Test Success!")
	}

}
