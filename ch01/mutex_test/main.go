package main

import (
	"fmt"
	"sync"
	"time"
)

var mx sync.Mutex

var (
	x = 10
	y = 0
)

func read(value int) {
	fmt.Println("Read: ", value)
}

func write(value *int) {
	*value = *value + 1
	fmt.Println("Write: ", *value)
}

func sum(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

func threadProc(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Thread %d: requesting lock\n", id)

	mx.Lock()
	defer mx.Unlock()

	fmt.Printf("Thread %d: lock acquired\n", id)

	// 임계 영역 시작
	read(x)
	write(&y)

	time.Sleep(500 * time.Millisecond)

	sum(x, y)
	// 임계 영역 끝

	fmt.Printf("Thread %d: releasing lock\n", id)
}

func main() {
	var wg sync.WaitGroup // 자체적으로 동시 접근을 안전하게 처리
	wg.Add(2)

	// 어느 goroutine이 먼저 잠금을 획득할지는 보장 X
	// 하지만 한쪽이 임계 영역을 실행하는 동안 다른 쪽은 같은 임계 영역에 진입 X
	go threadProc(1, &wg)
	go threadProc(2, &wg)

	wg.Wait()

	fmt.Println("Final y: ", y)

}
