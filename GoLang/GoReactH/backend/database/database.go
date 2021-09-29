package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type db_wraper struct {
	db *sql.DB
}

type TESTVAL struct {
	ID          string
	title       string
	description string
}

func Connect_to_db() (db_wraper, error) {
	var connection db_wraper
	var err error
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "PracticeB",
	}
	connection.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return connection, fmt.Errorf("%v; connecting to the database", err)
	}
	return connection, nil
}

func (conn *db_wraper) Add_test_values() error {
	db := conn.db
	_, err := db.Exec("DROP TABLE IF EXISTS test;")
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE test (
		id INT AUTO_INCREMENT NOT NULL,
		title VARCHAR(128) NOT NULL,
		description VARCHAR(255) NOT NULL
	);`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`INSERT INTO test (title, description)
	VALUES
	('First Title', 'First Description'),
	('Second Title', 'Second Description'),
	('Third Title', 'Third Description');
	`)
	if err != nil {
		return err
	}
	return nil
}

func (conn *db_wraper) Get_test_values() ([]TESTVAL, error) {
	db := conn.db
	var vals []TESTVAL

	rows, err := db.Query("SELECT * FROM test;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item TESTVAL
		if err := rows.Scan(&item.ID, &item.title, &item.description); err != nil {
			return nil, fmt.Errorf("GetTestValues: %v", err)
		}
		vals = append(vals, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetTestValues:: %v", err)
	}
	return vals, nil
}
