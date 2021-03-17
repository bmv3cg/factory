package main

import "fmt"

func main() {

	g := GetClient("gcp")
	//var obj ObjectStore

	g = GetClient("gcp")
	g.gcp.CreateBucket("gcp-bucket-15-3")
	fmt.Println(g.gcp.ListBuckets("diesel-harmony-306813"))

	c := GetClient("aws")
	c.aws.CreateBucket("aws-bukcet-test-15-3")
	c.aws.Listbuckets()
}
