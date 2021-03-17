package main

type ObjectStore interface {
	CreateBucket(string) error
	GetBucket(string) error
	ListBuckets(string) ([]string, error)
	DeleteBucket(string) error
}

type Client struct {
	aws AWSClient
	gcp GCPClient
}

func GetClient(cloudProvider string) Client {
	//var client ObjectStore
	switch cloudProvider {
	case "aws":
		return NewclientAWS("us-west-2")
	case "gcp":
		return NewClientGCP()
	default:
		return Client{}

	}

}
