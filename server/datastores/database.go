package datastores

import (
	"log"

	"github.com/jmoiron/sqlx"
)

// NewDBConnectionPool creates a new sqlx.DB connecting to DATABASE_URL
func NewDBConnectionPool(url string) *sqlx.DB {
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}
