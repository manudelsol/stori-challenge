package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"time"
)

func CreateAccount(ctx context.Context, email string, balance float64, txCount int64, destId *int64, tx pgx.Tx) error {
	err := tx.QueryRow(ctx, "INSERT INTO accounts (created_at, user_email, balance, number_of_txs) VALUES ($1, $2, $3, $4) RETURNING id",
		time.Now(), email, balance, txCount,
	).Scan(destId)

	if err != nil {
		return err
	}
	return nil
}
