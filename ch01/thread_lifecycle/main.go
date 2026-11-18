package main

import (
	"fmt"
	"sync" // 표준 동기화 패키지
	"time"
)

func threadProc(value int, wg *sync.WaitGroup) { // threadProc()와 main()이 같은 WaitGroup을 사용하도록 포인터 전달
	// 현재 함수가 반환되기 직전에 실행할 함수 호출 예약
	defer wg.Done() // 작업 1개 완료

	// 5. 다른 작업 처리
	fmt.Println("Worker started: ", value)
	time.Sleep(100 * time.Millisecond)

	// 6. 함수 실행 끝
	fmt.Println("Worker finished.") // 이후 defer 키워드로 예약된 함수 호출
}

func main() {
	var wg sync.WaitGroup // 등록된 작업이 모두 끝날 때까지 기다리는 도구
	wg.Add(1)             // 작업 1개 등록

	// 1. 새로운 goroutine 시작(생성 요청, 호출한 쪽은 함수가 종료될 때까지 기다리지 않고 다음 코드 진행)
	// 비동기 흐름을 만드는 것은 맞으나, 실행 환경에 따라 병렬 실행도 가능(실제 OS 스레드 배경은 런타임 스케줄러가)
	go threadProc(123, &wg) // Go의 함수 인수는 기본적으로 값 전달, 즉 값으로 전달하면 '복사'되므로 '동기화'가 안됨

	// 2. 메인에서도 작업 수행
	fmt.Println("Main working")
	time.Sleep(50 * time.Millisecond)

	// 3. goroutine 종료까지 대기
	wg.Wait() // 작업 카운터가 0이 될때까지 대기

	// 4. 함수 실행 끝
	fmt.Println("Main finished.")
}
