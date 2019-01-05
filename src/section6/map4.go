//자료형 : 맵(4)

package main

import "fmt"

func main() {
	//맵(Map)
	//맵 조회 할 경우에 주의할 점

	//예제1
	map1 := map[string]int{ //int : 0, string : "", float : 0.0
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}

	value1, ok1 := map1["lemon"]
	value2, ok2 := map1["kiwi"] //없는 키값으로 접근하는 경우 value는 값 타입의 초기화값을 넣어버림
	//따라서 없는 키값을 리턴 시 두번째 리턴 값으로 키 존재 유무 확인
	value3, ok3 := map1["kiwi"]

	fmt.Println("ex1 : ", value1, ok1)
	fmt.Println("ex1 : ", value2, ok2)
	fmt.Println("ex1 : ", value3, ok3)

	//예제2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : kiwi is not exist.")
	}

	if value, ok := map1["lemon"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : lemon is not exist.")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 : ", value)
	} else {
		fmt.Println("ex2 : banana is not exist.")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 : kiwi is not exist.")
	}

}
