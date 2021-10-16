// Copyright @ 2021 github.com/tianyagk
// Licensed under the GNU GENERAL PUBLIC LICENSE (the "License");
// you may not use this file except in compliance with the License.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// CliToolkit is a small cmd package for Go cmd shell interaction application
// inspired by go and cobra

package CliToolkit

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Event struct {
	DoFunc       func(s string, cmd Command) error
	Description  string
	Flag         string
	ErrorHandler func(e error)
}

type Command struct {
	Use     string
	Intro   string
	Short   string
	Long    string
	Prompt  string
	FuncMap map[string]Event
}

func (cmd Command) Run() {
	// Init do_help func
	cmd.FuncMap["help"] = Event{doHelp, "Cli command help", "-h", DefaultErrorHandler}
	cmd.FuncMap["exit"] = Event{doExit, "Exit Cli Toolkit", "-e", DefaultErrorHandler}

	// Print intro string
	fmt.Println(cmd.Intro)
	fmt.Println(cmd.Short)
	prompt := cmd.Prompt

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
		doExecute(cmdString, cmd)
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

func doExecute(cmdString string, cmd Command) {
	funcMap := cmd.FuncMap
	cmdString = strings.TrimSuffix(cmdString, "\n")
	arrCommandStr := strings.Fields(cmdString)
	for cmdStr, event := range funcMap {
		if (arrCommandStr[0] == cmdStr) || (arrCommandStr[0] == event.Flag) {
			if len(cmdString) == len(arrCommandStr[0]) {
				err := event.DoFunc(strings.TrimPrefix(cmdString, cmdString), cmd)
				if err != nil {
					event.ErrorHandler(err)
				}
				return
			} else {
				err := event.DoFunc(strings.TrimPrefix(cmdString, arrCommandStr[0]+" "), cmd)
				if err != nil {
					event.ErrorHandler(err)
				}
				return
			}
		}
	}
	// Not found command handler
	DefaultErrorHandler(errors.New("Can not find command: " + arrCommandStr[0]))
}

func doHelp(str string, cmd Command) error {
	if str == "" {
		for key, value := range cmd.FuncMap {
			fmt.Println(key, "	|Flag:", value.Flag, "	|", value.Description)
		}
	} else {
		if _, ok := cmd.FuncMap[str]; ok {
			fmt.Println(str, "	|Flag: ", cmd.FuncMap[str].Flag, "	| ", cmd.FuncMap[str].Description)
		} else {
			return errors.New("Can not find command: " + str)
		}
	}
	return nil
}

func doExit(str string, cmd Command) error {
	fmt.Println("Exit " + cmd.Use + str)
	os.Exit(0)
	return nil
}
