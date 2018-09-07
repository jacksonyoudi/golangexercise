package main

import "fmt"

//import "fmt"
//
//func ad(s []int)  ([]int){
//	s = append(s, 3)
//	return s
//}
//
//
//func main()  {
//	s := make([]int, 1, 2)
//	fmt.Println(s)
//	s = ad(s)
//	fmt.Println(s)
//}

//func main()  {
//	t := time.Now()
//	fmt.Println(t.Format("Mon Jan _2 15:04:05 2006"))
//}

func main()  {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}

	select {
	}
}