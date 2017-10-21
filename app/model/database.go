package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USER")
	dbname   = os.Getenv("DB_NAME")
	password = os.Getenv("DB_PASSWORD")
)

type Database struct {
	Conn *gorm.DB
}

func NewDatabase() (*Database, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, password))
	if err != nil {
		return nil, err
	}

	return &Database{
		Conn: db,
	}, nil
}

func (db *Database) SetupDB() {
	db.Conn.CreateTable(&Post{})

	db.Conn.AutoMigrate(&Post{})
}

