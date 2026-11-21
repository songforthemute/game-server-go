package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	maxCount    = 150000
	threadCount = 4
)

func isPrimeNumber(number int) bool {
	if number == 1 {
		return false
	}

	if number == 2 || number == 3 {
		return true
	}

	for i := 2; i < number-1; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func printNumber(primes []int) {
	for _, v := range primes {
		fmt.Println(v)
	}
}

func main() {
	// 모든 goroutine이 공유하는 변수
	num := 1
	var primes []int

	t0 := time.Now()

	var wg sync.WaitGroup
	wg.Add(threadCount)

	// 4개의 goroutine 생성
	for range threadCount {
		go func() {
			defer wg.Done()

			for {
				n := num
				num++

				if n >= maxCount {
					break
				}

				if isPrimeNumber(n) {
					primes = append(primes, n)
				}
			}
		}()
	}

	// 모든 goroutine의 종료를 대기
	wg.Wait()

	duration := time.Since(t0).Milliseconds()

	fmt.Printf("Took %d milliseconds. \n", duration)
}
