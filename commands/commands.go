package commands

import (
	"errors"
	"fmt"
	"github.com/FunkyLoiso/NotesProto/core"
	"github.com/FunkyLoiso/NotesProto/edit"
	"github.com/FunkyLoiso/NotesProto/list"
	"log"
	"os"
	"strings"
)

type commandInfo struct {
	Execute     func() error
	Usage       func()
	Description string
}

var cmds map[string]commandInfo

func init() {
	cmds = map[string]commandInfo{
		"list": {list.Exec, list.PrintUsage, "List notes from a notepad, newest first"},
		"edit": {edit.Exec, edit.PrintUsage, "Edit note or create new"},
		"help": {execHelp, printGeneralHelp, "Show help on specific command"},
	}
}

func execHelp() error {
	switch len(os.Args) {
	case 2:
		printGeneralHelp()
	case 3:
		if cmd, found := cmds[os.Args[2]]; !found {
			return errors.New("Unknown command '" + os.Args[2] + "'")
		} else {
			cmd.Usage()
		}
	default:
		return errors.New("Unknown command '" + strings.Join(os.Args[2:], " ") + "'")
	}
	return nil
}

func printGeneralHelp() {
	fmt.Printf("%v - somewhat potentially ok notes manager.\nCommands:\n", core.ExecName)

	maxCmdLength := 0
	for cmd := range cmds {
		if maxCmdLength < len(cmd) {
			maxCmdLength = len(cmd)
		}
	}
	for cmd, info := range cmds {
		fmt.Printf("%-*v%v\n", maxCmdLength+4, cmd, info.Description)
	}
	fmt.Printf("see '%v help <command>' for command details\n", core.ExecName)
}

func ParseAndExec() (err error) {
	var (
		cmd   commandInfo
		found bool
	)
	if len(os.Args) > 1 {
		cmd, found = cmds[os.Args[1]]
	} else {
		found = false
	}
	if !found {
		printGeneralHelp()
	} else {
		log.Printf("Executing comand '%v'", os.Args[1])
		err = cmd.Execute()
		if err != nil {
			log.Printf("Error while executing command '%v': %v", os.Args[1], err)
		}
	}
	return
}
