# CliToolkit

[![Go Report Card](https://goreportcard.com/badge/github.com/tianyagk/CliToolkit)](https://goreportcard.com/report/github.com/tianyagk/CliToolkit)  
**Blocking command line shell interaction tool.**

CliToolkit is a small cmd package for Go cmd shell interaction application.

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
CommandClient := CliToolkit.Command{
		Use:    "DefaultApp",
		Intro:  "CliToolkit Application",
		Short:  "Hello, welcome CliToolkit by Golang",
		Long:   "long:",
		Prompt: ">> ",
	}

FuncMap := make(map[string]CliToolkit.Event)
FuncMap["echo"] = CliToolkit.Event{DoFunc: echo, Description: "Repeat input string", Flag: "-echo", ErrorHandler: CliToolkit.DefaultErrorHandler}

CommandClient.FuncMap = FuncMap
CommandClient.Run()

// Define your command func here
func echo(str string, _ CliToolkit.Command) error {
	fmt.Println(str)
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



## Contributions
