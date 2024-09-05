package s3

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func ReadCSVFromS3(ctx context.Context, bucket, key string) (*csv.Reader, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
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

	csvReader := csv.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return csvReader, nil
}
