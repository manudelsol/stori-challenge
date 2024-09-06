package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"stori-challenge/src/db"
	"stori-challenge/src/email"
	"stori-challenge/src/s3"
	"stori-challenge/src/utils"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	Email  string `json:"email"`
}

func handleRequest(event *MyEvent) error {
	ctx := context.TODO()
	rows, err := s3.ReadCSVFromS3(ctx, event.Bucket, event.Key)
	if err != nil {
		fmt.Println(err)
		return err
	}

	conn, err := db.ConnectDB(ctx)
	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	varMap, balance, txCount, err := utils.ProcessRecords(rows[1:])

	destId := new(int64)
	err = db.CreateAccount(ctx, event.Email, balance, txCount, destId, tx)
	if err != nil {
		return err
	}

	err = db.CreateTxs(ctx, rows[1:], destId, tx)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	err = email.SendEmail(event.Email, varMap)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//Start the Lambda handler
	lambda.Start(handleRequest)
}
