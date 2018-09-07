package main

import "fmt"

//func main() {
//	c := make(chan bool, 0)
//go func() {
//	fmt.Println("GO GO GO!!!!!")
//	<-c
//}()
//c<-true
////for v := range c {
////	fmt.Println(v)
////}
//}

//func main() {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	c := make(chan bool, 10)
//	for i := 10; i < 10; i++ {
//		c <- true
//	}
//	for i := 0; i < 10; i++ {
//		go Go(c, i)
//	}
//}
//
//func Go(c chan bool, index int) {
//	a := 1
//	for i := 0; i < 100000000000; i++ {
//		a += i
//	}
//	fmt.Println(index, a)
//	<- c
//}

//func main()  {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	wg := sync.WaitGroup{}
//	wg.Add(10)
//
//	for i:=0;i<10;i++ {
//		go GO(&wg, i)
//	}
//
//	wg.Wait()
//}
//
//func GO(wg *sync.WaitGroup, index int)  {
//	a := 1
//	for i :=0;i < 10000000; i++ {
//		a += i
//	}
//	fmt.Println(index, a)
//
//	wg.Done()
//}

func main()  {
	c1,c2 := make(chan int), make(chan string)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <- c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	c1 <- 1
	c2 <- "hi"
	c1 <- 3
	c2 <- "hello"

	close(c1)
	close(c2)
	<- o

}





























