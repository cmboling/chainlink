package migrationsv2

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"io/ioutil"
	"path"
	"runtime"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "1_initial",
		Migrate: func(db *gorm.DB) error {
			_, filename, _, _ := runtime.Caller(0)
			query, err := ioutil.ReadFile(path.Join(path.Dir(filename), "1_initial/up.sql"))
			if err != nil {
				return err
			}
			return  db.Exec(string(query)).Error
		},
		Rollback: func(db *gorm.DB) error {
			_, filename, _, _ := runtime.Caller(0)
			query, err := ioutil.ReadFile(path.Join(path.Dir(filename), "1_initial/down.sql"))
			if err != nil {
				return err
			}
			return db.Exec(string(query)).Error
		},
	})
}