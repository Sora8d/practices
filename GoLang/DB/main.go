package main

import (
	"C"
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type db_wrap struct {
	db *sql.DB
}

var lite db_wrap

func main() {
	var err error
	lite.db, err = sql.Open("sqlite", "./experimeting.db")
	if err != nil {
		log.Fatal(err)
	}
	pingErr := lite.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Conexion done")

	err = lite.setupalbumtable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Seems your table has been correctly implemented")
}

func (dwrap *db_wrap) setupalbumtable() error {
	_, err := dwrap.db.Exec("DROP TABLE IF EXISTS album;")
	if err != nil {
		return fmt.Errorf("%v in dropping tables", err)
	}
	_, err = dwrap.db.Exec(`CREATE TABLE album (
		id	INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		title VARCHAR(128) NOT NULL,
		artist VARCHAR(255) NOT NULL,
		price DECIMAL(5,2) NOT NULL
	);`)
	if err != nil {
		return fmt.Errorf("%v in table creation", err)
	}
	_, err = dwrap.db.Exec(`
	INSERT INTO album 
  (title, artist, price) 
	VALUES 
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
	`)
	if err != nil {
		return fmt.Errorf("%v creating sample data", err)
	}
	return nil
}

func (dwrap *db_wrap) albumsByArtist(name string) ([]album, error) {
	var albums []album

	rows, err := dwrap.db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields
	for rows.Next() {
		var alb album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func (dwrap *db_wrap) albumByID(id int64) (album, error) {
	var alb album

	row := dwrap.db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func (dwrap *db_wrap) addAlbum(alb album) (int64, error) {
	result, err := dwrap.db.Exec("INSERT INTO album (title, artist, price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}
