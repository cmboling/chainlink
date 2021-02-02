package migrationsv2

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // http://doc.gorm.io/database.html#connecting-to-a-database
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMigrate_Migrations(t *testing.T) {
	db, err := gorm.Open("postgres", "postgres://postgres:node@localhost:5432/postgres?sslmode=disable")
	require.NoError(t, err)
	defer db.Close()

	dbname := "test"
	err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbname)).Error
	require.NoError(t, err)
	err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname)).Error
	require.NoError(t, err)
	db2, err := gorm.Open("postgres", "postgres://postgres:node@localhost:5432/test?sslmode=disable")
	require.NoError(t, err)
	err = MigrateUp(db2, "")
	require.NoError(t, err)
	err = db2.Exec("select * from bridge_types").Error
	require.NoError(t, err)
}
