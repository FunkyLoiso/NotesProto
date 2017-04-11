PRAGMA foreign_keys = off;
PRAGMA user_version = 100;
BEGIN TRANSACTION;

-- Table: Color
CREATE TABLE Color (ID INTEGER PRIMARY KEY UNIQUE NOT NULL, Name TEXT UNIQUE NOT NULL);
INSERT INTO Color (ID, Name) VALUES (1, 'white');
INSERT INTO Color (ID, Name) VALUES (2, 'green');
INSERT INTO Color (ID, Name) VALUES (3, 'teal');
INSERT INTO Color (ID, Name) VALUES (4, 'blue');
INSERT INTO Color (ID, Name) VALUES (5, 'gray');
INSERT INTO Color (ID, Name) VALUES (6, 'yellow');
INSERT INTO Color (ID, Name) VALUES (7, 'orange');
INSERT INTO Color (ID, Name) VALUES (8, 'red');

-- Table: Note
CREATE TABLE Note (ID INTEGER PRIMARY KEY UNIQUE NOT NULL, Notepad REFERENCES Notepad (ID) NOT NULL, Modified TEXT NOT NULL, Title TEXT NOT NULL, Text TEXT NOT NULL, Favorite INTEGER, Archived INTEGER);

-- Table: Notepad
CREATE TABLE Notepad (ID INTEGER PRIMARY KEY UNIQUE NOT NULL, Name TEXT UNIQUE NOT NULL, color REFERENCES Color (ID));
INSERT INTO Notepad (ID, Name, color) VALUES (1, 'Inbox', 1);

-- View: Recent
CREATE VIEW Recent AS SELECT
      nt.ID,
      nt.Favorite,
      nt.Title,
      nt.Text,
      np.Name Notepad,
      c.Name Color
  FROM Note nt
      JOIN
      Notepad np ON nt.Notepad = np.ID
      JOIN
      Color c ON np.color = c.ID
 WHERE nt.Archived IS NULL
 ORDER BY datetime(nt.Modified) DESC;

COMMIT TRANSACTION;
PRAGMA foreign_keys = on;