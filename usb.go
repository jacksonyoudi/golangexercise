package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}


type Connecter interface {
	Connect()
}


type Phone struct {
	name string
}

func (pc Phone) Name() string {
	return pc.name
}

func (pc Phone) Connect() {
	fmt.Println("Connect:", pc.name)
}


func main() {
	a := Phone{("youdi")}
	var b Connecter
	b = Connecter(a)
	b.Connect()
}


//func Disconnect(usb USB) {
//	if pc, ok := usb.(Phone); ok {
//		fmt.Println("disconnect", pc.name)
//		return
//	}
//	fmt.Println("Unknow device")
//}

func Disconnect(usb interface{}) {
	//if pc, ok := usb.(Phone); ok {
	//	fmt.Println("disconnect", pc.name)
	//	return
	//}

	switch v := usb.(type) {
	case Phone:
		fmt.Println("Disconnnet", v.name)
	default:
		fmt.Println("unkonw")

	}


	var a interface{}
	fmt.Println(a == nil)

	var p *int = nil
	fmt.Println(a == nil)



	//fmt.Println("Unknow device")
}

// 嵌入接口

// 接口转换
//可以降级

