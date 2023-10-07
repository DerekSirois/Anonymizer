package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	server := args[0]
	dbName := args[1]

	db, err := newDB(server, dbName)
	handleErrors(err)
	defer closeDb(db)

	persons, err := getAllPerson(db)
	handleErrors(err)

	for _, p := range persons {
		fmt.Printf("firstname: %s, lastname: %s, birthday: %v", p.Firstname, p.Lastname, p.Birthday)
	}
}

func handleErrors(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
