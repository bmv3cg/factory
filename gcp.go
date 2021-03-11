package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

type GCPclient struct {
	Client *storage.Client
	Ctx    context.Context
}

func NewGCPclient() *GCPclient {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("failed to create storage client")
	}
	return &GCPclient{Client: c, Ctx: ctx}
}

func (c *GCPclient) CreateBucketGCP(projectID, bucketName string) error {

	bucket := c.Client.Bucket(bucketName)
	if err := bucket.Create(c.Ctx, projectID, nil); err != nil {
		return fmt.Errorf("Bucket(%q).Create: %v", bucketName, err)
	}
	fmt.Println("Bucket created", bucketName)
	return nil
}

func (c *GCPclient) DeleteBucketGCP(bucketName string) error {

	bucket := c.Client.Bucket(bucketName)
	if err := bucket.Delete(c.Ctx); err != nil {
		return fmt.Errorf("Bucket(%q).Delete: %v", bucketName, err)
	}
	fmt.Println("Bucket %v deleted\n", bucketName)
	return nil
}

// listBuckets lists buckets in the project.
func (c *GCPclient) ListBucketsGCP(projectID string) ([]string, error) {

	var buckets []string
	it := c.Client.Buckets(c.Ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		buckets = append(buckets, battrs.Name)
		fmt.Println("Bucket:", battrs.Name)
	}
	return buckets, nil
}
