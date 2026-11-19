package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func FrameMove() {
	fmt.Println("FrameMove")
}

func Render() {
	fmt.Println("Render")
}

func load(name string) {
	fmt.Printf("Loading %s...\n", name)
	time.Sleep(500 * time.Millisecond)
}

func LoadScene()     { load("Scene") }
func LoadModel()     { load("Model") }
func LoadTexture()   { load("Texture") }
func LoadAnimation() { load("Animation") }
func LoadSound()     { load("Sound") }

var isStillLoading atomic.Bool

func threadProc1(wg *sync.WaitGroup) {
	defer wg.Done()

	for isStillLoading.Load() {
		FrameMove()
		Render()
		time.Sleep(200 * time.Millisecond)
	}
}

func threadProc2(wg *sync.WaitGroup) {
	defer wg.Done()

	LoadScene()
	LoadModel()
	LoadTexture()
	LoadAnimation()
	LoadSound()

	isStillLoading.Store(false)
}

func main() {
	isStillLoading.Store(true)

	var wg sync.WaitGroup
	wg.Add(2)

	go threadProc1(&wg)
	go threadProc2(&wg)

	wg.Wait()

	fmt.Println("Loading completed!")
}
