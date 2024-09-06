package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

// ConnectDB connects to the database
func ConnectDB(ctx context.Context) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection successful")
	return conn, nil
}
