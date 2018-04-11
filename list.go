package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) info() {
	fmt.Printf("name:%s, email:%s", u.name, u.email)
}

func (u *user) chang_email(email string) {
	u.email = email
}

func printInfo(u user) {
	fmt.Printf("name:%s, email:%s", u.name, u.email)
}

func main() {
	youdi := user{"youdi", "liangchangyou@gmail.com"}
	youdi.info()
	fmt.Printf("\n")
	youdi.chang_email("18684541304@163.com")
	youdi.info()
}
