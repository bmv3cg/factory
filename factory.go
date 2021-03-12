package main

type ObjectStore interface {
	Newclient(string)
	CreateBucket(string)
	GetBucket(string)
	Listbuckets(string)
	DeleteBucket(string)
}

func GetClient(cloudProvider string) ObjectStore {
	//var objstore ObjectStore
	switch cloudProvider {
	case "aws":
		return ObjectStore.Newclient("us-west-2")
	case "gcp":
		return ObjectStore.Newclient()
	default:
		return nil
	}

}
