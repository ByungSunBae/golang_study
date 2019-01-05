//자료형 : 포인터(1)

package main

import "fmt"

func main() {
	//포인터
	//Go : 포인터 지원(C)
	//변수의 지역성, 연속된 메모리 참조 ..., 힙, 스택
	//파이썬, 자바 (JRE) -> 컴파일러, 인터프리터
	//포인터지원X (파이썬, C#, JAVA 등)
	//주소의 값은 직접 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	//*(애스터리스크)로 사용
	//nil로 초기화 (nil != 0)

	//예제1
	var a *int            //방법1
	var b *int = new(int) //방법2

	fmt.Println(a) //&
	fmt.Println(b)

	i := 7

	a = &i // & 주소값 전달
	b = &i

	*a = 77

	fmt.Println("ex1 : ", a, &i)
	fmt.Println("ex1 : ", &a) //a는 현재 i의 주소값을 받음. 즉, 그 주소값을 저장한 메모리의 주소를 보려면 &a를 출력
	fmt.Println("ex1 : ", *a) //역참조, i의 주소값에 있는 값에 접근

	fmt.Println()

	fmt.Println("ex1 : ", b, &i)
	fmt.Println("ex1 : ", &b)
	fmt.Println("ex1 : ", *b) //역참조

	var c = &i
	d := &i

	*d = 10

	fmt.Println("ex1 : ", c, &i)
	fmt.Println("ex1 : ", &c)
	fmt.Println("ex1 : ", *c) //역참조, i의 주소값에 있는 값에 접근

	fmt.Println()

	fmt.Println("ex1 : ", d, &i)
	fmt.Println("ex1 : ", &d)
	fmt.Println("ex1 : ", *d) //역참조

}
