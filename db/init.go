package db

import (
	"strconv"
)

var buildVersion = 100

var initScript = `PRAGMA foreign_keys = off;
PRAGMA user_version = ` + strconv.Itoa(buildVersion) + `;

-- Table: Color
DROP TABLE IF EXISTS Color;
CREATE TABLE Color (id INTEGER PRIMARY KEY UNIQUE NOT NULL, name TEXT UNIQUE NOT NULL);
INSERT INTO Color (id, name) VALUES (1, 'white');
INSERT INTO Color (id, name) VALUES (2, 'green');
INSERT INTO Color (id, name) VALUES (3, 'teal');
INSERT INTO Color (id, name) VALUES (4, 'blue');
INSERT INTO Color (id, name) VALUES (5, 'gray');
INSERT INTO Color (id, name) VALUES (6, 'yellow');
INSERT INTO Color (id, name) VALUES (7, 'orange');
INSERT INTO Color (id, name) VALUES (8, 'red');

-- Table: Note
DROP TABLE IF EXISTS Note;
CREATE TABLE Note (id INTEGER PRIMARY KEY UNIQUE NOT NULL, notepadid REFERENCES Notepad (id) NOT NULL, modified DATETIME NOT NULL, title TEXT NOT NULL, text TEXT NOT NULL, favorite INTEGER NOT NULL DEFAULT 0, archived INTEGER NOT NULL DEFAULT 0);
INSERT INTO Note (id, notepadid, modified, title, text, favorite, archived) VALUES (15, 1, '2017-04-15T12:02:18+03:00', 'Destroy all humans', 'Destroy all humans
1. Use virus;
2. wait.', 0, 0);

-- Table: Notepad
DROP TABLE IF EXISTS Notepad;
CREATE TABLE Notepad (id INTEGER PRIMARY KEY UNIQUE NOT NULL, name TEXT UNIQUE NOT NULL, color REFERENCES Color (id));
INSERT INTO Notepad (id, name, color) VALUES (1, 'Inbox', 1);

-- View: NoteFull
DROP VIEW IF EXISTS NoteFull;
CREATE VIEW NoteFull AS SELECT
      nt.id,
      nt.notepadid,
      np.name notepadname,
      c.name notepadcolor,
      nt.modified,
      nt.title,
      nt.text,
      nt.favorite,
      nt.archived
  FROM Note nt
      JOIN
      Notepad np ON nt.notepadid = np.id
      JOIN
      Color c ON np.color = c.id;

-- View: Recent
DROP VIEW IF EXISTS Recent;
CREATE VIEW Recent AS SELECT
      nt.id,
      nt.favorite,
      nt.title,
      nt.text,
      np.name notepad,
      c.name color
  FROM Note nt
      JOIN
      Notepad np ON nt.notepadid = np.id
      JOIN
      Color c ON np.color = c.id
 WHERE nt.archived IS NULL
 ORDER BY datetime(nt.modified) DESC;


PRAGMA foreign_keys = on;`
