package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// var DB *sql.DB

func DbIn() (*sql.DB, error) {
	var err error
	connStr := `host=localhost port=5432 user=postgres dbname=jwt password=Pawan@2003 sslmode=disable`
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Could not ping database:", err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	// Uncomment this if you need to create tables
	createTable(DB)
	return DB, nil
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users(
	id TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL
)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
		return
	}
}
