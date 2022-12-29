package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
	} `yaml:"mysql"`
}

func loadConfig(fileName string) (Config) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error loading config file")
		panic(err)
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		fmt.Println("Error unmarshaling config yaml")
		panic(err)
	}

	return config
}

func connectDb(conf Config) (*sql.DB) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", conf.Mysql.User, conf.Mysql.Password, conf.Mysql.Address, conf.Mysql.Port))
	if err != nil {
		fmt.Println("Error connecting to mysql")
		panic(err)
	}
	return db
}

func getSchemas(db *sql.DB) ([]string) {
	return query(db, "SHOW DATABASES")
}

func getTablesBySchema(db *sql.DB, schemaName string) ([]string) {
	q := "SELECT TABLE_NAME FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = '"+schemaName+"'"
	return query(db, q)
}

func query(db *sql.DB, query string) ([]string) {
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("error querying db")
		panic(err)
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var s string
		if err := rows.Scan(&s); err != nil {
			fmt.Println("error iterating results")
			panic(err)
		}
		result = append(result, s)
	}
	return result
}
