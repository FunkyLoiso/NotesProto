package list

import (
	"flag"
	"fmt"
	"github.com/FunkyLoiso/NotesProto/core"
	"log"
)

var _flags *flag.FlagSet
var idStr string

func flags() *flag.FlagSet {
	if _flags != nil {
		return _flags
	}

	_flags = flag.NewFlagSet("list", flag.ContinueOnError)
	_flags.StringVar(&idStr, "n", "", "limit output to N most recent notes")
	_flags.Usage = PrintUsage
	return _flags
}

func PrintUsage() {
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
	log.Println("starting list command")

	return nil
}
