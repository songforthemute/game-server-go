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

func printNumbers(primes []int) {
	for _, v := range primes {
		fmt.Println(v)
	}
}

func main() {
	var (
		num      = 1
		numMutex sync.Mutex

		primes      []int
		primesMutex sync.Mutex // 뮤텍스 하나로 두 자원을 보호하면 서로 독립적인 작업인데도 기다려야 함
		// 잠금 세분성(Lock Granularity) : 뮤텍스가 보호하는 자원과 코드의 범위르 적절하게 설계하는 것

		t0 = time.Now()
		wg sync.WaitGroup
	)

	wg.Add(threadCount)

	for range threadCount {
		go func() {
			defer wg.Done()

			for {
				// 임계 영역 1: 다음 검사 대상 획득, 직렬 처리
				numMutex.Lock()
				n := num
				num++
				numMutex.Unlock()

				if n >= maxCount {
					break
				}

				// 소수 판별: 공유 자원에 접근 X, 병렬 처리
				if isPrimeNumber(n) {
					// 임계 영역 2: 결과 저장, 직렬 처리
					primesMutex.Lock()
					primes = append(primes, n)
					primesMutex.Unlock()
				}
			}
		}()
	}

	wg.Wait()

	duration := time.Since(t0).Milliseconds()

	fmt.Printf("Took %d milliseconds. \n", duration)

	// printNumbers(primes)
	fmt.Println("Prime count: ", len(primes))
}
