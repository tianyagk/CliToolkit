# CliToolkit

[![Go Report Card](https://goreportcard.com/badge/github.com/tianyagk/CliToolkit)](https://goreportcard.com/report/github.com/tianyagk/CliToolkit)  ![Lines of code](https://img.shields.io/tokei/lines/github/tianyagk/CliToolkit)  ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tianyagk/CliToolkit)  
**Blocking command line shell interaction tool.**

CliToolkit is a small cmd package for Go cmd shell interaction REPL(**Read-Eval-Print Loop**) application.

## Installation

```go
# download from github
go get github.com/tianyagk/CliToolkit
```

## Using the ClikToolkit Library

**Create CliToolkit**

Initial your interaction app with CliToolkit:

```go
// Init and Setup Command Client with Function Mapper
CommandClient := CliToolkit.Cli{
		Use:    "DefaultApp",
		Intro:  "CliToolkit Application",
		Short:  "Hello, welcome CliToolkit by Golang",
		Long:   "long:",
		Prompt: ">> ",
	}

FuncMap := make(map[string]CliToolkit.Event)
FuncMap["echo"] = CliToolkit.Event{DoFunc: echo, Description: "Repeat input string", ErrorHandler: CliToolkit.DefaultErrorHandler}

CommandClient.FuncMap = FuncMap
CommandClient.Run()

// Define your command func here
func echo(args []string, _ Cli) error {
	fmt.Println(args)
	return nil
}
```

You will additionally define the command function and ErrorHandler function into your FuncMap.

**Example custom function**

Define your own ErrorHandler function for individual command function:

```go
func CustomErrorHandler(err error) {
	if err != nil {
        fmt.Println("Command Error:", err)
	}
}
```



## Example

demo.go

```go
package main

import (
	"github.com/tianyagk/CliToolkit"
	"errors"
	"fmt"
)

func main() {
	g := make(map[string]int)
	g["init"] = 10

	_, ok := g["init"]
	fmt.Println(ok)

	CommandClient := CliToolkit.Cli{
		Use:    "DemoApp",
		Intro:  "CliToolkit Application",
		Short:  "Hello, welcome CliToolkit by Golang",
		Long:   "long:",
		Prompt: ">> ",
	}

	FuncMap := make(map[string]CliToolkit.Event)
	FuncMap["echo"] = CliToolkit.Event{DoFunc: echo, Description: "Repeat input string", CliToolkit.ErrorHandler: DefaultErrorHandler}
	FuncMap["error"] = CliToolkit.Event{DoFunc: errorMaker, Description: "Make an error", CliToolkit.ErrorHandler: DefaultErrorHandler}

	CommandClient.FuncMap = FuncMap
	CommandClient.Run()
}

func echo(args []string, _ Cli) error {
	fmt.Println(args)
	return nil
}

func errorMaker(args []string, _ Cli) error {
	return errors.New(fmt.Sprint("trouble maker ", args))
}

```

Launch in the command line:

```shell
CliToolkit Application
Hello, welcome CliToolkit by Golang
>> help
error 	| Make an error
help 	| Cli command help
exit 	| Exit Cli Toolkit
echo 	| Repeat input string
>> echo hello
hello
>> error Ivan
trouble maker Ivan
>> exit
```
