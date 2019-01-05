//채널(Channel) 기초(6)
package main

import (
	"fmt"
)

func main() {
	//채널(Channel)
	//Close : 채널 닫기

	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 77777
		}
	}()

	//두번째 값은 channel에서 값을 수신받았는지 여부를 할당한다.
	//즉, 이를 이용해 if문으로 프로그램의 제어가 가능하다.
	val1, ok1 := <-ch
	fmt.Println("ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("ex2 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("ex3 : ", val3, ok3)

	close(ch) //채널 닫기
	val4, ok4 := <-ch
	fmt.Println("ex4 : ", val4, ok4)
}
