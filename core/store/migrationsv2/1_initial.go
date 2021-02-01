package migrationsv2

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"io/ioutil"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "1_initial",
		Migrate: func(db *gorm.DB) error {
			query, err := ioutil.ReadFile("1_initial/up.sql")
			if err != nil {
				panic(err)
			}
			if err := db.Raw(string(query)).Error; err != nil {
				panic(err)
			}
			return nil
		},
		Rollback: func(db *gorm.DB) error {
			query, err := ioutil.ReadFile("1_initial/down.sql")
			if err != nil {
				panic(err)
			}
			return db.Exec(string(query)).Error
		},
	})
}