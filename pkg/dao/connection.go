package dao

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

func get() *sql.DB {
	config, err := GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func GetConfiguration() (Configuration, error) {
	config := Configuration{}
	file, err := os.Open("./configuration.json")
	if err != nil {
		return config, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
