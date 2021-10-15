package cli_toolkit

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var Conf = make(map[string]string)

type Event struct {
	DoFunc func(s string) error
	Description string
	Flag string
	ErrorHandler func(e error)
}
var FuncMap = map[string]Event{}

func Client(config map[string]string, funcMap map[string]Event) error{
	// Init do_help func
	funcMap["help"] = Event{doHelp, "Cli command help", "-h", DefaultErrorHandler}
	funcMap["exit"] = Event{doExit, "Exit Cli Toolkit", "-e", DefaultErrorHandler}

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
				return err
			}
		}
		// Execute Command
		doExecute(cmdString, funcMap)
	}
}

func DefaultErrorHandler(err error) {
	if err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return 
		}
	}
}

func doExecute(cmdString string, funcMap map[string]Event) {
	cmdString = strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(cmdString)
	for cmd, event := range funcMap {
		if (arrCommandStr[0] == cmd) || (arrCommandStr[0] == event.Flag) {
			if len(cmdString) == len(arrCommandStr[0]) {
				err := event.DoFunc(strings.TrimPrefix(cmdString, cmdString))
				if err != nil {
					event.ErrorHandler(err)
				}
				return
			} else {
				err := event.DoFunc(strings.TrimPrefix(cmdString, arrCommandStr[0]+" "))
				if err != nil {
					event.ErrorHandler(err)
				}
				return
			}
		}
	}
	// Not found command handler
	DefaultErrorHandler(errors.New("Can not find command: "+arrCommandStr[0]))
}


func doHelp(str string) error {
	if str == "" {
		for key, value:= range FuncMap {
			fmt.Println(key, "	|Flag:", value.Flag,"	|", value.Description)
		}
	} else {
		if _, ok := FuncMap[str]; ok {
			fmt.Println(str, "	|Flag: ", FuncMap[str].Flag,"	| ", FuncMap[str].Description)
		} else {
			return errors.New("Can not find command: "+str)
		}
	}
	return nil
}


func doExit(str string) error {
	fmt.Println("Exit "+Conf["name"] + str)
	os.Exit(0)
	return nil
}
