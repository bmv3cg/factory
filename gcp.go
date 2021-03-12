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

func (g *GCPclient) Newclient() string {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("failed to create storage client")
	}
	g = &GCPclient{Client: client, Ctx: ctx}
	return "GCP client created"
}

func (g *GCPclient) CreateBucket(projectID, bucketName string) error {

	bucket := g.Client.Bucket(bucketName)
	if err := bucket.Create(g.Ctx, projectID, nil); err != nil {
		return fmt.Errorf("Bucket(%q).Create: %v", bucketName, err)
	}
	fmt.Println("Bucket created", bucketName)
	return nil
}

func (g *GCPclient) DeleteBucket(bucketName string) error {

	bucket := g.Client.Bucket(bucketName)
	if err := bucket.Delete(g.Ctx); err != nil {
		return fmt.Errorf("Bucket(%q).Delete: %v", bucketName, err)
	}
	fmt.Println("Bucket %v deleted\n", bucketName)
	return nil
}

// listBuckets lists buckets in the project.
func (g *GCPclient) ListBuckets(projectID string) ([]string, error) {

	var buckets []string
	it := g.Client.Buckets(g.Ctx, projectID)
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
