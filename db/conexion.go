package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func GetConexion() (*sql.DB, error) {
	var user, pass, host, port, db string
	user = os.Getenv("USER")
	pass = os.Getenv("PASSWORD")
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
	db = os.Getenv("DATABASE")

	strconnect := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, pass, host, port, db)
	DB, _ = sql.Open("mysql", strconnect)
	// if err != nil {
	// 	return DB, err
	// }
	return DB, nil
}
