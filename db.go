package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type DBConfig struct {
	User     string `json:"db_user"`
	Password string `json:"db_password"`
	Host     string `json:"db_host"`
	Port     string `json:"db_port"`
	Name     string `json:"db_name"`
}

func LoadConfig(path string) (*DBConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg DBConfig
	err = json.NewDecoder(file).Decode(&cfg)
	return &cfg, err
}

func InitDB() (*sql.DB, error) {
	cfg, err := LoadConfig("config.json")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
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
