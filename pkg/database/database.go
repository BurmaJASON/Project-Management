package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Ensure the MySQL driver is imported
)

func Connect() (*sql.DB, error) {
	connStr := "user:userpassword@tcp(mysql:3306)/project_management"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database successfully")
	return db, nil
}
