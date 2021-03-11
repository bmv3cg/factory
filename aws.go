package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSClient struct {
	Client *s3.S3
}

func NewAWSclient(region string) *AWSClient {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	s3Client := s3.New(sess)
	return &AWSClient{Client: s3Client}
}

func (c *AWSClient) CreateBucketAWS(bucket string) error {

	_, err := c.Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	// Wait until bucket is created before finishing
	err = c.Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}
	return nil
}

// GetBucket determines whether we have this bucket
func (c *AWSClient) GetBucketAWS(bucket string) error {

	_, err := c.Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *AWSClient) ListbucketsAWS() (b []string) {

	result, err := c.Client.ListBuckets(nil)
	if err != nil {
		fmt.Println("Unable to list buckets")
		fmt.Println(err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
	return b
}

func (c *AWSClient) DeleteBucketAWS(bucket string) error {

	// Delete the S3 Bucket
	_, err := c.Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	// Wait until bucket is gone before finishing
	err = c.Client.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	_, err = c.Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil
	}

	return err

}
