package main

import (
	"CliToolkit/CliToolkit"
	"fmt"
)

func main() {
	CliToolkit.Conf["name"] = "demoCliApplication"
	CliToolkit.Conf["intro"] = "Launch CMD Toolkit."
	CliToolkit.Conf["short"] = "This toolkit for Golang CMD dev, entry .help for more information."
	CliToolkit.Conf["prompt"] = "$ "

	// 使用 struct 构建 FuncMap
	CliToolkit.FuncMap["echo"] = CliToolkit.Event{DoFunc: echo, Description: "Echo input", Flag: "-echo"}
	CliToolkit.Client(CliToolkit.Conf, CliToolkit.FuncMap)
}


func echo(str string) {
	fmt.Println(str)
}