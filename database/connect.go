package database

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/joho/godotenv"
	"os"
	
)


func ConnectDatabase() (DB *sql.DB, err error) {
	// Sambung ke Database
	err = godotenv.Load("config/dbpath.env")
	if err != nil {
		fmt.Println("failed to load config")
	} else {
		fmt.Println("success loaded config")
	}

	dbAddress := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("PGHOST"), 
	os.Getenv("PGPORT"), 
	os.Getenv("PGUSER"), 
	os.Getenv("PGPASSWORD"), 
	os.Getenv("PGDATABASE"))


	DB, err = sql.Open("postgres", dbAddress)
	err = DB.Ping()
	if err != nil {
		fmt.Println("Gagal menyambung ke Database")
		panic(err)
	} else {
		fmt.Println("database connection established")
	}

	return DB, err
}