package main

import (
		"gopkg.in/ini.v1"
)

func main() {
	cfg, _ := ini.LooseLoad("ga.ini")
	cfg.Reload()
	cfg, _ = ini.Load("ga.ini")

	//cfg.Section("youdi").Key("secret").SetValue("ypdxwomt2koos6h5")
	//cfg.SaveTo("ga.ini")
}