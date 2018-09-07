package main

import "fmt"

//type TZ int

//type A struct {
//	name string
//}

type In int

func main() {
	var i In
	i = 1
	i.Increase()
	fmt.Println(i)
	//var a TZ
	//a.Print()
	//(*TZ).Print(&a)
}

func (i *In) Increase() {
		*i = *i + 100
}


//func (a *TZ) Print() {
//	fmt.Println("TZ")
//}
//
//
//func (a *A) modify() {
//	a.name = "123"
//	fmt.Println(a.name)
//}

// method的访问权限是很高的
//struct的访问权限是在包级别的访问控制
//method value
   
//method Expression
