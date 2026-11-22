package main

import (
	"fmt"
	"sync"
)

/**
# go run -race ...
    - go가 제공하는 데이터 경쟁 탐지기(Race Detector) 활성화
	- go 컴파일러가 메모리 접근을 추적할 수 있도록 코드를 계측하고, 실행 중에 서로 다른 goroutine의 충돌하는 메모리 접근 감지
	- 계측으로 인해 실행 속도, 스케줄링 타이밍이 달라질 수 있음
*/

func main() {
	var (
		num     int = 0
		numbers []int
	)

	const (
		workerCount int = 4
		iterations      = 1000
	)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	for range workerCount {
		go func() {
			defer wg.Done()

			for range iterations {
				n := num
				num++

				numbers = append(numbers, n)
			}
		}()
	}

	wg.Wait()

	fmt.Println("Expected: ", workerCount*iterations)
	fmt.Println("Counter: ", num)
	fmt.Println("Slice length: ", len(numbers))
}
