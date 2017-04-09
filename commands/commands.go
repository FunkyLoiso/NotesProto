package commands

import (
	"github.com/FunkyLoiso/NotesProto/edit"
	"github.com/FunkyLoiso/NotesProto/list"
)

type CommandInfo struct {
	Execute     func() error
	Description string
}

var Commands = map[string]CommandInfo{
	"list": {list.Exec, "List notes from a notepad, newest first"},
	"edit": {edit.Exec, "Edit note or create new"},
	// "help": {help, "Show help on specific command"},
}
