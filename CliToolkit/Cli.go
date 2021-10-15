package CliToolkit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Conf = make(map[string]string)

type Event struct {
	DoFunc func(s string)
	Description string
	Flag string
}
var FuncMap = map[string]Event{}

func Client(config map[string]string, funcMap map[string]Event) {
	// Init do_help func
	funcMap["help"] = Event{doHelp, "Cli command help", "-h"}
	funcMap["exit"] = Event{doExit, "Exit Cli Toolkit", "-e"}

	// Print intro string
	intro, ok := config["intro"]
	if ok {
		fmt.Println(intro)
	}

	short, ok := config["short"]
	if ok {
		fmt.Println(short)
	}

	// Set prompt style
	prompt, ok := config["prompt"]
	if !ok {
		prompt = ">> "
	}

	// Blocking interaction
	for {
		fmt.Print(prompt)
		reader := bufio.NewReader(os.Stdin)
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return 
			}
		}
		// Execute Command
		doExecute(cmdString, funcMap)
	}
}

func doExecute(cmdString string, funcMap map[string]Event) {
	cmdString = strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(cmdString)
	for idx, value := range funcMap {
		if (arrCommandStr[0] == idx) || (arrCommandStr[0] == value.Flag) {
			if len(cmdString) == len(arrCommandStr[0]) {
				value.DoFunc(strings.TrimPrefix(cmdString, cmdString))
				return
			} else {
				value.DoFunc(strings.TrimPrefix(cmdString, arrCommandStr[0]+" "))
				return
			}
		}
	}
	// Not found command handler
	fmt.Println("Not find command: "+arrCommandStr[0])
}

func doHelp(str string) {
	if str == "" {
		for key, value:= range FuncMap {
			fmt.Println(key, "	|Flag:", value.Flag,"	|", value.Description)
		}
	} else {
		if _, ok := FuncMap[str]; ok {
			fmt.Println(str, "	|Flag: ", FuncMap[str].Flag,"	| ", FuncMap[str].Description)
		} else {
			fmt.Println("Can not find command: ", str)
		}
	}
}


func doExit(str string) {
	fmt.Println("Exit "+Conf["name"])
	os.Exit(0)
}
