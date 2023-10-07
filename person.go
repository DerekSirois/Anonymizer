package main

import (
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/lib/pq"
)

type Person struct {
	Id        int       `fake:"skip"`
	Firstname string    `fake:"{firstname}"`
	Lastname  string    `fake:"{lastname}"`
	Birthday  time.Time `db:"birthday"`
}

func getAllPerson(db *sqlx.DB) ([]*Person, error) {
	var person []*Person
	err := db.Select(&person, "SELECT * from person")
	return person, err
}

func (p *Person) update(db *sqlx.DB) error {
	_, err := db.Exec("UPDATE person SET firstname = $1, lastname = $2, birthday = $3 WHERE id = $4", p.Firstname, p.Lastname, p.Birthday, p.Id)
	return err
}
