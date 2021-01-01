package infrastructure

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func init() {
	dbConfig := os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":3306)/" + os.Getenv("DATABASE_NAME") + "?parseTime=true"

	db, err := sqlx.Connect("mysql", dbConfig)
	if err != nil {
		panic(err.Error())
	}

	DB = db
}
