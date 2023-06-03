package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// benchmarkFileRead("heavyjob", func() { heavyJob("heavyjob") })

	// 1. 900ms
	// benchmarkFileRead("heavyjob more 3", func() {
	// 	heavyJob("heavy1")
	// 	heavyJob("heavy2")
	// 	heavyJob("heavy3")
	// })

	// 2. 600ms (1개 비동기)
	// 다만 main이 먼저 종료되면 heavy2번의 작업을 마치지 못한채로 프로그램이 종료됨
	// benchmarkFileRead("heavyjob more 3", func() {
	// 	heavyJob("heavy1")
	// 	go heavyJob("heavy2")
	// 	heavyJob("heavy3")
	// })

	// 3. 위와 동일 -> main 프로그램이 먼저 종료됨 sleep 함수를 사용? -> X
	// benchmarkFileRead("heavyjob more 3", func() {
	// 	go heavyJob("heavy1")
	// 	go heavyJob("heavy2")
	// 	go heavyJob("heavy3")
	// })

	// 4. use wait group (41ms)
	// var wg sync.WaitGroup
	// wg.Add(3)
	// benchmarkFileRead("heavyjob more 3", func() {
	// 	go heavyJobUseWg("heavy1", &wg)
	// 	go heavyJobUseWg("heavy2", &wg)
	// 	go heavyJobUseWg("heavy3", &wg)
	// })
	// wg.Wait()

	resultCh := make(chan int)
	benchmarkFileRead("heavyjob more 3", func() {
		wg := sync.WaitGroup{}
		wg.Add(3)
		go heavyJobUseChan("heavy1", resultCh, &wg)
		go heavyJobUseChan("heavy2", resultCh, &wg)
		go heavyJobUseChan("heavy3", resultCh, &wg)
		wg.Wait()
	})

	for i := 0; i < 3; i++ {
		<-resultCh
	}

}

func heavyJob(jobname string) {

	sum := 0
	for i := 0; i < 1000000000; i++ {
		// job
		sum += i
	}

	fmt.Printf("%s >> %d\n", jobname, sum)
}

func heavyJob2(jobname string, ch chan<- int) {

	sum := 0
	for i := 0; i < 1000000000; i++ {
		// job
		sum += i
	}

	fmt.Printf("%s >> %d\n", jobname, sum)
	ch <- sum
}

func heavyJobUseWg(jobname string, wg *sync.WaitGroup) {
	defer wg.Done()
	heavyJob(jobname)
}

func heavyJobUseChan(jobname string, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	heavyJob(jobname)
}

func benchmarkFileRead(text string, f func()) {
	start := time.Now()
	f()
	fmt.Println(text, " >> ", time.Since(start))
}
