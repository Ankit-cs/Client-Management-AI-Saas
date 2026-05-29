package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool" //to establish connection with postgresSql
)
//global context helps in timeouts and cancellation of database operations and closing any connections
func NewPool(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {

	pool,err:=pgxpool.New(ctx,databaseURL)
	if err!=nil{
		return nil,fmt.Errorf("unable to create connection pool: %v",err)
	}
	if err:=pool.Ping(ctx);err!=nil{
		pool.Close()
		return nil,fmt.Errorf("unable to connect to database: %v",err)
	}
	return pool,nil

}