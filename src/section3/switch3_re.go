//Switch문(3)

package main

import "fmt"

func main() {
	//예제1
	a := 30 / 15
	switch a {
	case 2, 4, 6: //a가 2, 4, 6인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5: //a가 1, 3, 5인 경우
		if a >= 20 {
			fmt.Println("20 이상")
		}
		fmt.Println("a -> ", a, "는 홀수")
	}

	//예제2
	//fallthrough : 조건에 맞지않아도 진입해서 연산 수행
	//그래서 case의 맨마지막 조건에는 넣을 수 없음
	switch e := "go"; e {
	case "java":
		fmt.Println("java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
	case "ruby":
		fmt.Println("ruby!")
	case "c":
		fmt.Println("c!")
	}
}
