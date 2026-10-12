package main

import (
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type S3Config struct {
	Region   string
	Bucket   string
	Endpoint string
}

func LoadS3Config() (*S3Config, error) {
	return &S3Config{
		Region:   os.Getenv("AWS_REGION"),
		Bucket:   os.Getenv("S3_BUCKET"),
		Endpoint: os.Getenv("S3_ENDPOINT"),
	}, nil
}

func (c *S3Config) Validate() error {
	if c.Bucket == "" {
		return aws.NewError("S3_BUCKET environment variable is required")
	}
	return nil
}
