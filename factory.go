package main

type ObjectStore interface {
	Newclient(string)
	CreateBucket(string)
	GetBucket(string)
	Listbuckets(string)
	DeleteBucket(string)
}

func GetClient(cloudProvider string) string {
	//var objstore ObjectStore
	switch cloudProvider {
	case "aws":
		return AWSClient{Client: nil}.Newclient("")
	case "gcp":
		return GCPclient{Client: nil, Ctx: nil}.Newclient()
	default:
		return ""
	}

}
