package db

import (
	"errors"
	"fmt"
	"github.com/FunkyLoiso/NotesProto/core"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var _ = fmt.Print
var _ = os.Open

const (
	defDbPath = "./db.sqlite3"
)

var db *sqlx.DB
var dbVersion int

// do nothing if already opened
func Open() error {
	if db != nil {
		return nil
	}

	dbPath := defDbPath
	if core.Cfg.DBPath != "" {
		dbPath = core.Cfg.DBPath
	}
	log.Printf("Opening db for path '%v'..\n", dbPath)
	var err error
	db, err = sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		log.Printf("sqlx.Connect failed for path '%v': %v\n", dbPath, err)
		return err
	}

	// check db version
	ver := 0
	err = db.Get(&ver, "PRAGMA user_version")
	if err != nil {
		log.Println("Failed to get user_version from db: ", err)
		return err
	}
	if ver == 0 {
		// db is empty
		err = createDb()
		if err != nil {
			return err
		}
		err = db.Get(&ver, "PRAGMA user_version")
		if err != nil {
			log.Println("Failed to get user_version after creating db: ", err)
			return err
		}
	}

	if ver/100 != buildVersion/100 {
		errStr := fmt.Sprintf("Incompatible db versions. application: %v, db: %v", buildVersion, ver)
		log.Printf(errStr)
		return errors.New(errStr)
	}
	dbVersion = ver
	return nil
}

func Close() {
	if db == nil {
		return
	}

	err := db.Close()
	if err != nil {
		log.Println("db.Close failed: ", err)
	}
	db = nil
}

func createDb() error {
	log.Println("Creating db..")
	tx, err := db.Beginx()
	if err != nil {
		log.Println("Failed to start transaction:", err)
		return err
	}
	r, err := tx.Exec(initScript)
	if err != nil {
		log.Println("Failed to execute db init script: ", err)
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("Failed to commit transaction: ", err)
		return err
	}
	affected, err := r.RowsAffected()
	if err != nil {
		log.Println("RowsAffected failed: ", err)
	}
	log.Println("DB created, rows affected: ", affected)
	return nil
}
