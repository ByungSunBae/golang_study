//변수1
package main

import "fmt"

func main() {
	//기본 초기화
	//정수타입 : 0, 실수(소수점) : 0.0, 문자열 : "", Boolean : true, false
	//변수명 : 숫자 첫글자x, 대소문자 구분O, 문자, 숫자, 밑줄, 특수기호 사용가능
	//변수 및 상수 : 함수 내외 사용가능
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j string = "Hi! Golang!"
	var k = 4.74 //선언 동시 초기화
	var l = "Hi! Seoul!"
	var m = true

	//나중에 선언해줘도 돌아간다.
	a = 4
	b = "Hello Go!"
	e = 77

	// 선언된 변수를 사용하지 않으면 Golang에서는 에러를 뱉음
	// 따라서 선언된 변수는 아래와 같이 반드시 사용을 해줘야함.
	// 꼭 print가 아니어도 되지만 반드시 사용해줄것!!

	fmt.Println("a : ", a)
	fmt.Println("b : ", b)
	fmt.Println("c : ", c)
	fmt.Println("d : ", d)
	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)
}
