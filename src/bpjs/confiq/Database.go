package confiq

import (
	"bpjs/model"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
)

func Connect() (db *sql.DB) {
	var config model.Config
	file, err := os.Open("D:/bpjs/src/bpjs/confiq/confiq.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}