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

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Event struct {
	DoFunc       func(args []string, cmd Cli) error
	Description  string
	Args         []string
	ErrorHandler func(e error)
}

type Cli struct {
	Use     string
	Intro   string
	Short   string
	Long    string
	Prompt  string
	FuncMap map[string]Event
}

func (cmd Cli) Run() {
	// Init do_help func
	cmd.FuncMap["help"] = Event{doHelp, "Cli command help", nil,DefaultErrorHandler}
	cmd.FuncMap["exit"] = Event{doExit, "Exit Cli Toolkit", nil,DefaultErrorHandler}

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
		// Execute Cli
		doExecute(cmdString, cmd)
	}
}

func DefaultErrorHandler(err error) {
	if err != nil {
		fmt.Print("\r")
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
	}
}

func doExecute(cmdString string, cmd Cli) {
	funcMap := cmd.FuncMap

	// remove mark '\n', '\r'
	cmdString = strings.TrimSuffix(cmdString, "\n")
	cmdString = strings.TrimSuffix(cmdString, "\r")
	arrCommandStr := strings.Fields(cmdString)

	flag := arrCommandStr[0]
	doEvent, ok := funcMap[flag]
	if !ok{
		DefaultErrorHandler(errors.New("Can not find command: " + arrCommandStr[0]))
	} else {
		args := strings.Fields(strings.TrimPrefix(cmdString, arrCommandStr[0]))
		err := doEvent.DoFunc(args, cmd)
		if err != nil {
			doEvent.ErrorHandler(err)
		}
	}
}

func doHelp(args []string, cmd Cli) error {
	if len(args) == 0 {
		for key, value := range cmd.FuncMap {
			fmt.Println(key, "	|", value.Description)
		}
	} else {
		if _, ok := cmd.FuncMap[args[0]]; ok {
			fmt.Println(args, "	|Flag: ", "	| ", cmd.FuncMap[args[0]].Description)
		} else {
			return errors.New("Can not find command: " + args[0])
		}
	}
	return nil
}

func doExit(args []string, cmd Cli) error {
	fmt.Println("Exit " + cmd.Use + args[0])
	os.Exit(0)
	return nil
}
