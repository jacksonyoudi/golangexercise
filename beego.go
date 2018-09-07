package main

import "github.com/astaxie/beego"

type HomeControler struct {
	beego.Controller
}

func (this *HomeControler) Get() {
	this.Ctx.WriteString("hello world!!!")
}

func main()  {
	beego.Router("/", &HomeControler{})
	beego.Run()
}