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
}
