package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	Client *sql.DB
)

func init() {
	//TODO: change to environment variable - now just to learn go
	dsn := fmt.Sprintf(
		"host='%s' port='%s' user='%s' password='%s' dbname='%s' sslmode='%s'",
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"bookstore",
		"disable",
	)

	var err error
	Client, err = sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

}
