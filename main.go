package main

import "fmt"

func main() {

	//var obj ObjectStore
	//
	//obj.CreateBucket("sadss")

	g := GetClient("gcp")
	g.gcp.CreateBucket("gcp-bucket-15-3")
	fmt.Println(g.gcp.ListBuckets("diesel-harmony-306813"))

	c := GetClient("aws")
	c.aws.CreateBucket("aws-bukcet-test-15-3")
	c.aws.Listbuckets()

}
