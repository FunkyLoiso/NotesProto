package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Notes proto be here")

	lfile, err := os.OpenFile("./NotesProto.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file")
	} else {
		log.SetOutput(lfile)
		log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	}

	log.Println("==================== NotesProto new log entry ====================")

	var conf config
	err = conf.read()
	fmt.Println(err)

	conf.Source = "local"

	err = conf.write()
	fmt.Println(err)
	log.Println("NotesProto stop")
}
