package configs

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/savinnsk/api-template-go/internal/infra/sqlc"
)


var dbConn *sql.DB
type DbInstance struct { 
	db *sql.DB
	*sqlc.Queries
	} 

func InitDb(cfg *conf) *sql.DB {
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s",cfg.DBUser,cfg.DBPassword,cfg.DBHost,cfg.DBDatabase)
	var err error
	
	dbConn, err := sql.Open("mysql",dns)
	if  err != nil {
		panic(err)
	}

	return dbConn
}

func GetInstanceDB() *DbInstance {
	if dbConn == nil {
		log.Fatal("Database instance error.")
	}
	return &DbInstance{
		db : dbConn,
		Queries: sqlc.New(dbConn),
	}
}

func (dbInst *DbInstance) CallTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := dbInst.db.BeginTx(ctx,nil)
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			 log.Printf("error on rollback: %v, original error: %v", errRb,err)
			 return fmt.Errorf("error on rollback: %v, original error: %v", errRb,err)
		}
		return err
	}

	return tx.Commit()
}


