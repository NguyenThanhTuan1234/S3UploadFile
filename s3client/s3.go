package s3client

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3PutObject struct {
	s3Client manager.UploadAPIClient
}

type PutObject interface {
	Upload([]byte, string, string) error
}

func S3New(cfg aws.Config) PutObject {
	return &s3PutObject{
		s3Client: s3.NewFromConfig(cfg),
	}
}

func (s *s3PutObject) Upload(buffer []byte, fileName, bucket string) error {
	putObjectInput := s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buffer),
	}
	_, err := s.s3Client.PutObject(context.Background(), &putObjectInput)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
