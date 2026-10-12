package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Uploader(ctx context.Context, cfg *S3Config) (*S3Uploader, error) {
	cfg_aws, err := config.LoadDefaultConfig(ctx, config.WithRegion(cfg.Region))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg_aws)
	return &S3Uploader{
		client: client,
		bucket: cfg.Bucket,
	}, nil
}

func (u *S3Uploader) Close() error {
	return nil
}
