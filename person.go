package main

import (
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/lib/pq"
)

type Person struct {
	firstname string    `db:"firstname"`
	lastname  string    `db:"lastname"`
	birthday  time.Time `db:"birthday"`
}

func getAllPerson(db *sqlx.DB) ([]*Person, error) {
	var person []*Person
	err := db.Select(&person, "SELECT firstname, lastname, birthday from person")
	return person, err
}
