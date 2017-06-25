package main

import (
	"fmt"
	"github.com/FunkyLoiso/NotesProto/commands"
	"github.com/FunkyLoiso/NotesProto/core"
	"log"
	"os"
	"path"
)

func main() {
	// open log
	_, err := os.OpenFile("./NotesProto.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file")
	} else {
		// log.SetOutput(lFile)
	}

	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("==================== NotesProto new log entry ====================")
	log.Printf("starting with args: %v\n", os.Args[1:])

	// determine executable name
	{
		appPath, err := os.Executable()
		if err != nil {
			log.Printf("os.Executable failed: %v", err)
			appPath = os.Args[0]
		}
		core.ExecName = path.Base(appPath)
	}

	// read config
	err = core.Cfg.Read()
	if err != nil {
		log.Println("Error reading config file:", err)
	}

	if err = commands.ParseAndExec(); err != nil {
		fmt.Println(err)
	}
	log.Println("NotesProto stop")
}
