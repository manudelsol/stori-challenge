package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

type TxData struct {
	Email string
	Rows  [][]string
}

func InsertData(ctx context.Context, data TxData) error {
	conn, err := ConnectDB(ctx)
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	for _, row := range data.Rows {
		parsedDate, err := time.Parse("1/2", row[1])
		if err != nil {
			return err
		}
		createdAt := parsedDate.AddDate(2024, 0, 0)
		_, err = tx.Exec(ctx, "INSERT INTO txns (created_at, tx, user_email) VALUES ($1, $2, $3)", createdAt, row[2], data.Email)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
