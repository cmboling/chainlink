package migrationsv2

import (
	"fmt"
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

var Migrations []*gormigrate.Migration

func Migrate(db *gorm.DB) error {
	return MigrateUp(db, "")
}

func MigrateUp(db *gorm.DB, to string) error {
	g := gormigrate.New(db, &gormigrate.Options{
		//TableName:                 "migrations",
		UseTransaction:            false,
		ValidateUnknownMigrations: false,
	}, Migrations)

	if to == "" {
		to = Migrations[len(Migrations) - 1].ID
	}
	fmt.Println("migrating to", to)
	if err := g.MigrateTo(to); err != nil {
		return err
	}
	return nil
}

func MigrateDown(db *gorm.DB) error {
	g := gormigrate.New(db, nil, Migrations)
	return g.RollbackLast()
}

func Reset(db *gorm.DB) error {
	g := gormigrate.New(db, nil, Migrations)
	first := Migrations[0].ID
	if err := g.RollbackTo(first); err != nil {
		return err
	}
	return g.RollbackMigration(Migrations[0])
}

func Version(db *gorm.DB) string {
	// Find the latest migration which ran
	for i := len(Migrations) - 1; i >= 0; i-- {
		migration := Migrations[i]

		var m gormigrate.Migration
		err := db.Find(&m).Where("id = ?", migration.ID).Error
		if err != nil {
			panic(err)
		}
		if m.ID != "" {
			return migration.ID
		}
	}
	return ""
}
