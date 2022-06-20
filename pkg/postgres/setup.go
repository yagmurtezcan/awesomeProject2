package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Database struct {
	Db *gorm.DB
}

func NewDatabase() *Database {
	database, err := gorm.Open(postgres.Open("postgres://postgres:Aa!12@localhost:5432/awesomeProject2"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	return &Database{Db: database}
}

func (d *Database) AutoMigrate(models ...interface{}) {
	err := d.Db.AutoMigrate(models...)
	if err != nil {
		log.Fatalln("Migration error")
	}
}
