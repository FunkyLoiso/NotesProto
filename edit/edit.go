package edit

import (
	"flag"
	"fmt"
	"github.com/FunkyLoiso/NotesProto/core"
	"os"
)

var _ = fmt.Printf

var _flags *flag.FlagSet
var id string

func flags() *flag.FlagSet {
	if _flags != nil {
		return _flags
	}

	_flags = flag.NewFlagSet("edit", flag.ContinueOnError)
	_flags.StringVar(&id, "i", "", "id of the note to edit")
	_flags.Usage = PrintUsage
	return _flags
}

func PrintUsage() {
	// fmt.Println("Edit note or create new")
	fmt.Printf("Usage: %v edit [<title> | -i <id>]\n", core.ExecName)
	fmt.Println("Options:")
	flags().VisitAll(func(f *flag.Flag) {
		fmt.Printf("    -%-5v%v", f.Name, f.Usage)
		if f.DefValue != "" {
			fmt.Printf(" (%v)", f.DefValue)
		}
		fmt.Println()
	})
}

func Exec() error {
	err := flags().Parse(os.Args[2:])
	if err != nil {
		return err
	}

	if id != "" {
		// edit by id
	} else {
		// edit by title or create new
	}

	return nil
}
