package core

import (
	"github.com/FunkyLoiso/NotesProto/list"
)

type CommandInfo struct {
	Execute     func() error
	Description string
}

var Commands = map[string]CommandInfo{
	"list": {list.Exec, "List notes from a notepad, newest first"},
	// "help": {help, "Show help on specific command"},
}
