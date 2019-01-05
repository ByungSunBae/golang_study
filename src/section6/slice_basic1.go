//자료형 : 슬라이스(1)

package main

import "fmt"

func main() {
	//길이x(가변) -> 동적으로 크기가 늘어난다. 레퍼런스(참조 값) 타입
	//슬라이스 (길이 & 용량) 크기가 동적으로 할당가능

	//2가지 선언 방법 1. 배열처럼 선언, 2. make 함수 : make(자료형, 길이, 용량(생략시 길이))

	//예제1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	//slice2[0] = 1 //길이가 가변형이니까 완벽하게 값을 넣지 않고 인덱스에 해당하는 원소를 수정할 수 없음.
	slice3[4] = 10 //값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	//예제2
	var slice5 []int = make([]int, 5, 10)
	//길이가 5, 총용량은 10인 slice임. 10만큼 메모리를 확보함.
	//10보다 더 큰 인덱스의 값을 넣으면 그만큼의 메모리를 재할당함.
	//데이터가 큰 경우 지정한 총용량보다 크면 재할당이 일어나면서 그만큼 성능 차이가 일어남.
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)

	slice6[2] = 7

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)

	//예제3
	var slice9 []int //nil 슬라이스(길이와 용량 0), 선언만 되어있고 초기화 안되어있음. 재할당 불가!

	if slice9 == nil {
		fmt.Println("This is Nli Slice!")
	}
}
