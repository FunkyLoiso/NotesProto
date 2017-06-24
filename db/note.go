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
		log.Printf("Failed to get Note with id = '%x': %v", id, err)
		return nil, err
	}
	return &note, nil
}

func UpdateNote(n *Note) error {
	err := Open()
	if err != nil {
		return err
	}

	_, err = db.Exec(`
    UPDATE Note
      SET notepadid = ?, title = ?, text = ?, favorite = ?, archived = ?
      WHERE id = ?`,
		n.NotepadId, n.Title, n.Text, n.Favorite, n.Archived, n.Id)
	return err
}

func CreateNote(n *Note) error {
	err := Open()
	if err != nil {
		return err
	}

	r, err := db.Exec(`
    INSERT INTO Note (notepadid, title, text, favorite, archived)
      VALUES (?, ?, ?, ?, ?)`,
		n.NotepadId, n.Title, n.Text, n.Favorite, n.Archived)
	if err == nil {
		if id, err := r.LastInsertId(); err != nil {
			return err
		} else {
			n.Id = id
		}
	}
	return err
}
