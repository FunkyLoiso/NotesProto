package edit

import (
	"flag"
	"fmt"
	"github.com/FunkyLoiso/NotesProto/core"
	"github.com/FunkyLoiso/NotesProto/db"
	"log"
	"os"
	"strconv"
)

var _ = fmt.Printf

var _flags *flag.FlagSet
var idStr string

func flags() *flag.FlagSet {
	if _flags != nil {
		return _flags
	}

	_flags = flag.NewFlagSet("edit", flag.ContinueOnError)
	_flags.StringVar(&idStr, "i", "", "id of the note to edit")
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

	if idStr != "" {
		id, err := strconv.ParseInt(idStr, 16, 64)
		if err != nil {
			log.Printf("Failed to parse id string '%v' to int: %v", idStr, err)
			return err
		}
		// edit by id
		log.Printf("Retrieving note with id '%v'", id)
		note, err := db.GetNote(id)
		if err != nil {
			log.Printf("Failed to get note with id = '%v' from db: %v", id, err)
			return err
		}

		log.Println("Retrived note: ", *note)
		//
	} else {
		// edit by title or create new
	}

	return nil
}
