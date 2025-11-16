package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	s3Client   *s3.Client
	bucketName string
	awsRegion  string
)

// InitS3 initializes the S3 client and bucket configuration
func InitS3() {
	// 1. Get region and bucket name from environment
	awsRegion = os.Getenv("AWS_REGION")
	bucketName = os.Getenv("S3_BUCKET_NAME")

	if awsRegion == "" || bucketName == "" {
		log.Fatalf("AWS_REGION and S3_BUCKET_NAME must be set in .env")
	}

	// 2. Load the default config, which will find AWS_ACCESS_KEY_ID
	//    and AWS_SECRET_ACCESS_KEY from the environment
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(awsRegion))
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	s3Client = s3.NewFromConfig(cfg)
	log.Println("S3 Storage Service Initialized")
}

// UploadFile uploads a file to S3 and returns its public URL
// (Capital 'U' makes it exported)
func UploadFile(filename string, file io.Reader) (string, error) {
	_, err := s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		// 3. Use the bucketName variable
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    "public-read", // Makes the file publicly accessible
	})
	if err != nil {
		return "", err
	}

	// 4. Construct the public URL
	//    Note: S3 URL formats can vary. This is a common one.
	//    For "us-east-1", the region is sometimes omitted.
	//    A more robust way is to get this from the S3 SDK, but this
	//    is a standard and simple way.
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, awsRegion, filename)

	// Handle us-east-1 region URL format (which often has no region in the URL)
	if awsRegion == "us-east-1" {
		url = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, filename)
	}

	return url, nil
}
