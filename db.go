package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDB(server, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=dev password=abcde dbname=%s", server, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
