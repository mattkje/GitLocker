package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (*sql.DB, error) {
	// Example: root:password@tcp(127.0.0.1:3306)/git_system
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root",       // username
		"Ma10erkul",  // password
		"127.0.0.1",  // host
		"3306",       // port
		"git_system", // database name
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INT AUTO_INCREMENT PRIMARY KEY,
            username VARCHAR(255) NOT NULL UNIQUE,
            password TEXT NOT NULL
        )
    `)
	return db, err
}
