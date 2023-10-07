package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE if not exists person
(
    id SERIAL PRIMARY KEY,
	firstname text,
	lastname text,
	birthday date
);
`

func newDB(server, dbName string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=dev password=abcde dbname=%s sslmode=disable", server, dbName)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	db.MustExec(schema)
	return db, nil
}
