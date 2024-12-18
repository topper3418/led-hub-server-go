package db

import (
    "database/sql"
    "github.com/mattn/go-sqlite3"
    "log"
)

var conn *sql.DB

func init() {
    conn, err := sql.Open("sqlite3", "data/appData.db")
    if err != nil {
        log.fatalf("error getting connection to database %v", err)
    }
}
func getConnection() {
    return conn
}
