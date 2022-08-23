package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
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
func listEKSClusters() {
	svc := eks.New(session.New())
	input := &eks.ListClustersInput{}

	result, err := svc.ListClusters(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case eks.ErrCodeInvalidParameterException:
				fmt.Println(eks.ErrCodeInvalidParameterException, aerr.Error())
			case eks.ErrCodeClientException:
				fmt.Println(eks.ErrCodeClientException, aerr.Error())
			case eks.ErrCodeServerException:
				fmt.Println(eks.ErrCodeServerException, aerr.Error())
			case eks.ErrCodeServiceUnavailableException:
				fmt.Println(eks.ErrCodeServiceUnavailableException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func main() {

	/*myBucketName := "mybucket-" + (xid.New().String())

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load configuration")
	}

	s3client := s3.NewFromConfig(cfg)

	AccountBucketOps(*s3client, myBucketName)*/
	listEKSClusters()
	cmd.Execute()
}
