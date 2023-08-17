package dbDns

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const user string = "root"
const pass string = ""
const dbName string = "golang_db"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, pass, dbName))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
