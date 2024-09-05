package main

import (
	"context"
	"fmt"
	"stori-challenge/src/s3"
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
	fmt.Println(rows)

	return nil
}

func main() {
	// Start the Lambda handler
	//lambda.Start(handleRequest)
	handleRequest(&MyEvent{
		Bucket: "manubucket-demo-s3",
		Key:    "txns.csv",
		Email:  "",
	})
}
