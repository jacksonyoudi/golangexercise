package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)
	wg.Add(2)

	go playball("youdi", court)
	go playball("test", court)
	court <- 1

	wg.Wait()


}

func playball(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Println("player win", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Println("player missed", name)
			fmt.Println(n)
			close(court)
			return
		}

		fmt.Println("continue", name)
		ball++
		court <- ball

	}

}
