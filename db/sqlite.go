package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const (
	driverName   = "sqlite3"
	metadataPath = "./db/sqlite.db"
)

type SQLite struct {
	*sqlx.DB
}

func Init() SQLite {
	db, err := sqlx.Open(driverName, metadataPath)
	if err != nil {
		log.Println(err)
	}
	createBaseTables(db)
	sqlite := SQLite{DB: db}
	return sqlite
}

func createBaseTables(db *sqlx.DB) {
	_, err := db.Exec(createDailySignTable)
	if err != nil {
		panic(err)
	}
}
