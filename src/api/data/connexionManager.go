package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "gernika54"
	dbname   = "CRUD"
)

// GetConnection -> Init the connection to the database. /**
func GetConnection() *gorm.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")
	return db
}
