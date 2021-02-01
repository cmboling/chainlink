package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/smartcontractkit/chainlink/core/store/migrationsv2"
	"log"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres" // http://doc.gorm.io/database.html#connecting-to-a-database
)

func main() {
	// Connect to the database
	db, err := gorm.Open("postgres", os.Args[1])
	if err != nil {
		log.Fatal("unable to connect to db", err)
	}
	defer db.Close()

	switch os.Args[2] {
	case "up":
		migrationsv2.MigrateUp(db, os.Args[3])
	case "down":
		migrationsv2.MigrateDown(db)
	case "reset":
		migrationsv2.Reset(db)
	case "version":
		fmt.Println(migrationsv2.Version(db))
	default:
		log.Fatal("invalid use")
	}
}