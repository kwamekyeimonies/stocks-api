package database

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/kwamekyeimonies/stocks-api/models"
)

func Create_connection() *pg.DB {
	opts := &pg.Options{
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDRESS"),
		Database: os.Getenv("DB_DATABASE"),
	}

	db := pg.Connect(opts)
	if db == nil {
		log.Printf("Database Connection failed....\n")
		os.Exit(100)
		// return &pg.DB{}, true

	}

	log.Println("Database connected Successfully....")

	if err := createSchema(db); err != nil {
		log.Fatal(err)
	}

	return db

}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Stock)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
