package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	dbUser     = "passman"
	dbPassword = "passman"
	dbName     = "passman"
)

var schemaName = "public" //"passman_test"

var conn *sql.DB = nil

func getConnection() (*sql.DB, error) {
	var err error
	if conn == nil {
		connString := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=require search_path=%s", dbUser, dbPassword, dbName, 5432, schemaName)
		conn, err = sql.Open("postgres", connString)
		if err == nil {
			conn.SetMaxOpenConns(95)
			conn.SetMaxIdleConns(5)
		}
	}
	//fmt.Println(conn.Stats())
	return conn, err
}

func useTestSchema() {
	schemaName = "passman_test"
}
