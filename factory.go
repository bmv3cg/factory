package main

type ObjectStore interface {
	Newclient(string)
	CreateBucket(string)
	GetBucket(string)
	ListBuckets(string)
	DeleteBucket(string)
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
