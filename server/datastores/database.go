package datastores

import (
	"database/sql"
	"log"
)

// NewDBConnectionPool creates a new sql.DB connecting to DATABASE_URL
func NewDBConnectionPool(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
