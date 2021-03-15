package main

import "fmt"

func main() {

	//var obj ObjectStore
	//
	//obj.CreateBucket("sadss")

	//g := GetClient("gcp")
	// gcp.CreateBucket("ssssssfdsabcjhbaskjbf")
	//fmt.Println(g.gcp.ListBuckets("diesel-harmony-306813"))

	c := GetClient("AWS")
	fmt.Println(c)
	//c.aws.CreateBucket("sasf")
	//c.aws.Listbuckets()

}
