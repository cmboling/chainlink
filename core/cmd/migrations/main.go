package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // http://doc.gorm.io/database.html#connecting-to-a-database
	"github.com/smartcontractkit/chainlink/core/store/migrationsv2"
)

func main() {
	// TODO: Error handling, put in the client
	db, err := gorm.Open("postgres", os.Args[1])
	if err != nil {
		log.Fatal("unable to connect to db", err)
	}
	defer db.Close()

	switch os.Args[2] {
	case "up":
		if err := migrationsv2.MigrateUp(db, os.Args[3]); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := migrationsv2.MigrateDown(db); err != nil {
			log.Fatal(err)
		}
	case "reset":
		if err := migrationsv2.Reset(db); err != nil {
			log.Fatal(err)
		}
	case "version":
		fmt.Println(migrationsv2.Version(db))
	default:
		log.Fatal("invalid use")
	}
}
