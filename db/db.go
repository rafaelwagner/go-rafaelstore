package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func GetVariable(variable string) string {
	return os.Getenv(variable)
}

func DatabaseConector() *sql.DB {

	user := GetVariable("POSTGRES_USER")
	dbname := GetVariable("DB_NAME")
	port, _ := strconv.Atoi(GetVariable("DB_PORT"))
	password := GetVariable("POSTGRES_PASSWORD")
	host := GetVariable("DB_HOST")

	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
