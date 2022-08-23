package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/platform9/cloud-analyser/cmd"
)

func AccountBucketOps(client s3.Client, name string) {

	fmt.Println("List buckets: ")
	//snippet-start:[s3.go-v2.ListBuckets]
	listBucketsResult, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	if err != nil {
		panic("Couldn't list buckets")
	}
	for _, bucket := range listBucketsResult.Buckets {
		fmt.Printf("Bucket name: %s\t\tcreated at: %v\n", *bucket.Name, bucket.CreationDate)
	}

	if err != nil {
		panic("Couldn't list bucket contents")
	}

}

func main() {

	/*myBucketName := "mybucket-" + (xid.New().String())

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load configuration")
	}

	s3client := s3.NewFromConfig(cfg)

	AccountBucketOps(*s3client, myBucketName)*/
	cmd.Execute()
}
