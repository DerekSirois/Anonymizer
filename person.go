package main

import (
	"gorm.io/gorm"
	"time"
)

type Person struct {
	firstname string    `gorm:" column: firstname "`
	lastname  string    `gorm:" column: lastname "`
	birthday  time.Time `gorm:" column: birthday "`
}

func getAllPerson(db *gorm.DB) ([]Person, error) {
	var p []Person
	result := db.Table("person").Find(&p)
	return p, result.Error
}
