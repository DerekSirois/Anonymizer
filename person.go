package main

import (
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/lib/pq"
)

type Person struct {
	id        int       `db:"id"`
	firstname string    `db:"firstname"`
	lastname  string    `db:"lastname"`
	birthday  time.Time `db:"birthday"`
}

func getAllPerson(db *sqlx.DB) ([]*Person, error) {
	p := make([]*Person, 0)
	err := db.Select(&p, "SELECT * from person")
	return p, err
}
