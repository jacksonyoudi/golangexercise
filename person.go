package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) { //传递的是指针，直接修改原始的值
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1-struct as a value type:
	var pers1 Person // 类型
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	fmt.Println(pers1)
	upPerson(&pers1) //传递的是类型的指针
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2—struct as a pointer:
	pers2 := new(Person) //返回的是指针
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	fmt.Println(pers2) //　指针

	(*pers2).lastName = "Woodward" // 这是合法的
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	// 3—struct as a literal:
	pers3 := &Person{"Chris", "Woodward"} //返回的是指针　等同于　new(type)
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
