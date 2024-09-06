package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"time"
)

func CreateTxs(ctx context.Context, rows [][]string, accountId *int64, tx pgx.Tx) error {
	//txs creation
	for _, row := range rows {
		parsedDate, err := time.Parse("1/2", row[1])
		if err != nil {
			return err
		}
		createdAt := parsedDate.AddDate(2024, 0, 0)
		_, err = tx.Exec(ctx, "INSERT INTO txns (created_at, tx, account_id) VALUES ($1, $2, $3)", createdAt, row[2], *accountId)
		if err != nil {
			return err
		}
	}
	return nil
}
