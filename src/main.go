package main

import (
	"fmt"
)

type MyEvent struct {
	Bucket *string `json:"bucket"`
	Key    *string `json:"key"`
	Email  *string `json:"email"`
}

func handleRequest(event *MyEvent) error {
	fmt.Sprint(event)
	return nil
}

func main() {
	// Start the Lambda handler
	//lambda.Start(handleRequest)
}
