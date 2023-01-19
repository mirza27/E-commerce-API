package database

import(
	"database/sql"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DbConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB){
	migration := &migrate.PackrMigrationSource{
		Box: packr.New("migration", "./sql_migration"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migration, migrate.Up)
	if errs != nil{
		panic(errs)
	}

	DbConnection = dbParam

	fmt.Println("Telah melakukan ",n, " migrasi ke database")
}


