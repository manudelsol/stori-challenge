package s3

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ReadCSVFromS3(ctx context.Context, bucket, key string) ([][]string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	cfg.Region = "us-east-1"
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}
	svc := s3.NewFromConfig(cfg)
	resp, err := svc.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, fmt.Errorf("unable to download item, %v", err)
	}
	defer resp.Body.Close()

	rows, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
