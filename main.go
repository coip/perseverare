package main

import (
	"fmt"
	"os"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//needs: username, password, host, default_schema
const dsnform = "%s:%s@tcp(%s:3306)/%s"

var dsn string = fmt.Sprintf(dsnform,
	os.Getenv("username"),
	os.Getenv("password"),
	os.Getenv("host"),
	os.Getenv("default_schema"),
)

func init() {
	os.Stderr.Write([]byte("validating mariadb feasability...[" + dsn + "]\n"))
}

func main() {
	var (
		db  *sql.DB
		err error
	)

	if db, err = sql.Open("mysql", dsn); err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	os.Exit(0)

}
