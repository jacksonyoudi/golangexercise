package main

import (
	"reflect"
	"fmt"
	"bytes"
)

type Student struct {
	Name string
	Age int
}

type Class struct {
	Name string `json:"name"`
	Student *Student `json:"student"`
	Grade int `json:"grade"`
	School string `json:"school"`
}

func main() {

	//s := Student{Name: "LiLei", Age: 20}
	//
	//typ := reflect.TypeOf(s)
	//val := reflect.ValueOf(s)
	//
	//
	//fmt.Printf("%v", typ)
	//
	//fmt.Printf("The type is %s.\n", typ.String())
	//fmt.Printf("The name is %s.\n", val.FieldByName("Name").String())
	//
	//if s, ok := val.Interface().(Student); ok {
	//	fmt.Printf("The student is %s.\n", s.Name)
	//} else {
	//	fmt.Println("Wrong!")
	//}


	s := &Student{"LiLei", 18}
	c := &Class{"Class A", s, 6, "Century Ave"}

	val := reflect.ValueOf(c)
	typ := reflect.TypeOf(c)
	if val.Kind() == reflect.Ptr {
		fmt.Printf("It is a pointer. Address its value.\n")
		val = val.Elem()
		typ = typ.Elem()
	}

	for i := 0; i < val.NumField(); i = i + 1 {
		fv := val.Field(i)
		ft := typ.Field(i)
		switch fv.Kind() {
		case reflect.String:
			fmt.Printf("The %d th %s types %s valuing %s with tag env %s\n", i, ft.Name, "string", fv.String(), ft.Tag.Get("env"))
		case reflect.Int:
			fmt.Printf("The %d th %s types %s valuing %d with tag env %s\n", i, ft.Name, "int", fv.Int(), ft.Tag.Get("env"))
		case reflect.Ptr:
			fmt.Printf("The %d th %s types %s valuing %v with tag env %s\n", i, ft.Name, "pointer", fv.Pointer(), ft.Tag.Get("env"))
		}
	}


	var b bytes.Buffer

	for {
		str := "hello"
		b.WriteString(str)
	}


}
