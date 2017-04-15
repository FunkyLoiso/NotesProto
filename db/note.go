package db

import (
	"log"
	"time"
)

type Note struct {
	Id           int64
	NotepadId    int64
	NotepadName  string
	NotepadColor string
	Modified     time.Time
	Title        string
	Text         string
	Favorite     bool
	Archived     bool
}

func GetNote(id int64) (*Note, error) {
	err := Open()
	if err != nil {
		return nil, err
	}

	note := Note{}
	err = db.Get(&note, "SELECT * from NoteFull WHERE id = ?", id)
	if err != nil {
		log.Printf("Failed to get Note with id = '%v': %v", id, err)
		return nil, err
	}
	return &note, nil
}
