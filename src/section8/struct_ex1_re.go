//구조체 심화(1)
package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func NewAccount(number string, balance float64, interest float64) *Account { //포인터 반환 아닌 경우 값 복사
	return &Account{number, balance, interest} //구조체 인스턴스를 생성한 뒤 포인터를 리턴
}

func main() {
	//구조체 생성자 패턴 예제

	//예제1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.015}

	var lee *Account = new(Account)
	lee.number = "245-902"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1 : ", kim)
	fmt.Println("ex1 : ", lee)

	//예제2
	park := NewAccount("245-903", 17000000, 0.04)
	fmt.Println("ex2 : ", park)

	//예제3
	//에러 발생
	//NewAccount()라는 함수의 파라미터가 없기 때문
	// -> 이는 구조체 생성자는 Go에서 default value의 초기화를 하고싶지않은 경우 사용
	// -> 예제1에서 Account{}라고만 생성하면 각 필드의 값은 타입에 따라 default value로 초기화 수행
	// -> 만약 default value로 인해서 값이 잘못계산되면 여러가지 일들이 발생하기 때문이다.
	// bae := NewAccount()
	// fmt.Println("ex3 : ", bae)
}
