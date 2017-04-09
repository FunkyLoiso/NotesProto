package main

import (
	"fmt"
	"github.com/FunkyLoiso/NotesProto/core"
	"log"
	"os"
	"path"
)

func printHelp() {
	appPath, err := os.Executable()
	if err != nil {
		log.Printf("os.Executable failed: %v", err)
		appPath = os.Args[0]
	}
	execName := path.Base(appPath)
	fmt.Printf("%v - somewhat potentially ok notes manager.\nCommands:\n", execName)

	maxCmdLengh := 0
	for cmd, _ := range core.Commands {
		if maxCmdLengh < len(cmd) {
			maxCmdLengh = len(cmd)
		}
	}
	for cmd, info := range core.Commands {
		fmt.Printf("%-*v%v\n", maxCmdLengh+4, cmd, info.Description)
	}
	fmt.Printf("see '%v help <command>' for command details\n", execName)
}

func main() {
	// open log
	lFile, err := os.OpenFile("./NotesProto.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file")
	} else {
		log.SetOutput(lFile)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	}

	log.Println("==================== NotesProto new log entry ====================")
	log.Printf("starting with args: %v\n", os.Args)

	// read config
	err = core.Cfg.Read()
	if err != nil {
		log.Println("Error reading config file: %v", err)
	}

	// parse and execute command
	var (
		cmd   core.CommandInfo
		found bool
	)
	if len(os.Args) > 1 {
		cmd, found = core.Commands[os.Args[1]]
	} else {
		found = false
	}
	if !found {
		printHelp()
	} else {
		err = cmd.Execute()
		if err != nil {
			log.Printf("Error while executing command '%v': %v", os.Args[1], err)
		}
	}

	log.Println("NotesProto stop")
}
