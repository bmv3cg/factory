package main

type ObjectStoreClient interface{
	NewAWSclient(string)
	NewGCPclient()	
  }
  
  type ObjectStore interface{
	
	NewAWSclient(string)
	CreateBucketAWS(string)
	GetBucketAWS(string)
	ListbucketsAWS(string)
	DeleteBucketAWS(string)
	
	NewGCPclient()	
	CreateBucketGCP(string, string)
	ListBucketsGCP(string)
	DeleteBucketGCP(string)
  
  }
  
  
func GetClient(cloudProvider string)(ObjectStoreClient){
  
	switch cloudProvider {
	case "aws":
		return NewAWSclient("us-west-2")
	case "gcp":
	  return NewGCPclient()	  
	default: 
		return nil          
	}
  
}
  