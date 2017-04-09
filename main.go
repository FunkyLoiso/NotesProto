package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

var cfg config

type commandInfo struct {
	execute     func() error
	description string
}

var commands = map[string]commandInfo{
	"list": {list, "List notes from a notepad, newest first"},
}

func printHelp() {
	appPath, err := os.Executable()
	if err != nil {
		log.Printf("os.Executable failed: %v", err)
		appPath = os.Args[0]
	}
	execName := path.Base(appPath)
	fmt.Printf("%v - somewhat potentially ok notes manager.\nCommands:\n", execName)

	maxCmdLengh := 0
	for cmd, _ := range commands {
		if maxCmdLengh < len(cmd) {
			maxCmdLengh = len(cmd)
		}
	}
	for cmd, info := range commands {
		fmt.Printf("%-*v%v\n", maxCmdLengh+4, cmd, info.description)
	}
	fmt.Printf("see '%v help <command>' for command details\n", execName)
}

func main() {
	lFile, err := os.OpenFile("./NotesProto.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file")
	} else {
		log.SetOutput(lFile)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	}

	log.Println("==================== NotesProto new log entry ====================")
	log.Printf("starting with args: %v\n", os.Args)

	err = cfg.read()
	if err != nil {
		log.Println("Error reading config file: %v", err)
	}

	var (
		cmd   commandInfo
		found bool
	)
	if len(os.Args) > 1 {
		cmd, found = commands[os.Args[1]]
	} else {
		found = false
	}
	if !found {
		printHelp()
	} else {
		err = cmd.execute()
		if err != nil {
			log.Printf("Error while executing command '%v': %v", os.Args[1], err)
		}
	}

	log.Println("NotesProto stop")
}
