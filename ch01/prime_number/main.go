package main

import (
	"fmt"
	"time"
)

const maxCount = 150000

func isPrimeNumber(number int) bool {
	if number == 1 {
		return false
	}

	if number == 2 || number == 3 {
		return false
	}

	for i := 2; i < number-1; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func printNumbers(primes []int) {
	for _, v := range primes { // index, value
		fmt.Println(v)
	}
}

func main() {
	var primes []int // 슬라이스 타입, 동적 길이 시퀀스, 내부적으로 배열을 참조하는 구조

	t0 := time.Now()

	for i := 0; i <= maxCount; i++ {
		if isPrimeNumber(i) {
			primes = append(primes, i) // 요소를 추가한 결과 슬라이스 반환 - 내부 저장 공간이 부족하면 새로운 배열을 할당할 수도 있음
		}
	}

	t1 := time.Now()

	duration := t1.Sub(t0).Milliseconds()

	fmt.Printf("Took %d milliseconds. \n", duration)

	//printNumbers(primes)
}
