package edit

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/FunkyLoiso/NotesProto/core"
	"github.com/FunkyLoiso/NotesProto/db"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
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

		// open temp file for editor
		tempFile, err := ioutil.TempFile("", "nptemp_")
		if err != nil {
			log.Println("Failed to open temp file:", err)
			return err
		}
		defer os.Remove(tempFile.Name())

		// write text
		if _, err := tempFile.WriteString(note.Text); err != nil {
			log.Println("Failed to write note's content to temp file for editing:", err)
			return err
		}

		// open editor
		editor := core.Cfg.Editor
		if editor == "" {
			editor = core.GetDefaultEditor()
		}
		if editor == "" {
			errStr := fmt.Sprintln("Editor is undefined")
			log.Printf(errStr)
			return errors.New(errStr)
		}

		// editor might have some arguments so we try to separate them
		args := strings.Fields(editor)
		name := args[0]
		args = append(args[1:], tempFile.Name())

		cmd := exec.Command(name, args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		var editorStderr bytes.Buffer
		cmd.Stderr = &editorStderr
		log.Println("Going to run", cmd)
		if err = cmd.Run(); err != nil { // blocks until done
			log.Printf("Editor '%v' completed with error: %v", cmd.Path, err)
			if editorStderr.Len() != 0 {
				log.Println("Stderr is:", editorStderr.String())
			}
			return err
		}

		// determine if file changed
		if _, err := tempFile.Seek(0, 0); err != nil {
			log.Println("Seek to start failed for temp file:", err)
			return err
		}
		newBytes, err := ioutil.ReadAll(tempFile)
		if err != nil {
			log.Println("Failed to read temp file after editing:", err)
			return err
		}
		newText := string(newBytes)
		if note.Text != newText {
			log.Println("Text chaged to:\n", newText)
			// calculate title
			note.Title = core.MakeTitle(newText)
			note.Text = newText
			// add to db
			err = db.UpdateNote(note)
			if err != nil {
				log.Println("Failed to update note:", err)
				return err
			}
			log.Println("Note updated")

		} else {
			log.Println("Text is the same")
		}
		//
	} else {
		// edit by title or create new
	}

	return nil
}
