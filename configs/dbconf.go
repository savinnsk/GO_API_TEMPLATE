package configs

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var dbConn *sql.DB 

func InitDb(cfg *conf) *sql.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s",cfg.DBUser,cfg.DBPassword,cfg.DBHost,cfg.DBDatabase)
	var err error
	dbConn, err = sql.Open("mysql",dns)
	if  err != nil {
		panic(err)
	}

	return dbConn
}

func GetInstanceDB() *sql.DB {
	if dbConn == nil {
		log.Fatal("Database instance error.")
	}
	return dbConn
}