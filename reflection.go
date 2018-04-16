package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

//func (u user) Hello() {
//	fmt.Println("Hello world.")
//}

//type Manager struct {
//	User
//	title string
//}

func main() {
	u := User{1, "ok", 12}

	Set(&u)
	fmt.Println(u)

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")
	if !mv.IsValid() {
		fmt.Print("BAD")
		return
	}

	args := []reflect.Value{reflect.ValueOf("joe")}
	mv.Call(args)

	//x := 123
	//v := reflect.ValueOf(&x)
	//v.Elem().SetInt(999)
	//
	//fmt.Println(x)

	//m := Manager{User: User{1, "youdi", 123}, title: "youdi"}
	//t := reflect.TypeOf(m)

	//fmt.Printf("%#v\n", t.Field(1))
	//fmt.Printf("%#v\n", t.FieldByIndex([]int{1}))

	//u := User{1, "youdi", 12}
	//Info(u)
}

//func Info(o interface{}) {
//	t := reflect.TypeOf(o)
//	fmt.Println("type:", t.Name())
//
//	if k := t.Kind(); k != reflect.Struct {
//		fmt.Println("xxx")
//		return
//	}
//
//	v := reflect.ValueOf(o)
//	fmt.Println("fields:")
//
//	for i := 0; i < t.NumField(); i++ {
//		f := t.Field(i)
//		val := v.Field(i).Interface()
//		fmt.Printf("%s: %v = %v", f.Name, f.Type, val)
//	}
//
//	for i := 0; i < t.NumMethod(); i++ {
//		m := t.Method(i)
//		fmt.Print("%s: %v\n", m.Name, m.Type)
//	}
//
//}

func (u User) Hello(name string) {
	fmt.Println("hello", name, "my name is", u.Name)
}



func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr &&  !v.Elem().CanSet() {
		fmt.Print("xxx")
	} else {
		v = v.Elem()
	}

	if f := v.FieldByName("Name");f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}

}