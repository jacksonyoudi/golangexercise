package main

import (
	"sync"
	"sync/atomic"
	"runtime"
	"fmt"
)

var (
	counter int64
	wg sync.WaitGroup
	mutx sync.Mutex
)


func main(){
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final counter: ", counter)


}


func incCounter(id int) {
	defer wg.Done()


	mutx.Lock()
	for count := 0; count < 2;count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()

	}

	mutx.Unlock()
}