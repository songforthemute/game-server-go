package main

import "fmt"

func func1() {
	func2()
}

func func2() {
	func3()
}

func func3() { // 반환값이 없으면 타입 자체를 작성 X
	a := 0 // go는 타입 추론 지원, 선언한 지역 변수 미사용 시 컴파일 오류
	fmt.Println(a)
}

func main() {
	func1()
}
