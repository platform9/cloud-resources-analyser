package logic

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ryanuber/columnize"
)

func AccountBucketOps() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load configuration")
	}
	s3client := s3.NewFromConfig(cfg)
	input := &s3.ListBucketsInput{}

	fmt.Println("List buckets: ")

	//snippet-start:[s3.go-v2.ListBuckets]
	listBucketsResult, err := s3client.ListBuckets(context.TODO(), input)

	if err != nil {
		panic("Couldn't list buckets")
	}

	if err != nil {
		panic("Couldn't list bucket contents")
	}
	var Output []string
	var tableFormat = "Bucket name | created at: "
	Output = append(Output, tableFormat)
	for _, bucket := range listBucketsResult.Buckets {
		instanceInfo := fmt.Sprintf("%v | %v ", *bucket.Name, bucket.CreationDate)
		Output = append(Output, instanceInfo)
	}
	tabularClusterInfo := columnize.SimpleFormat(Output)
	fmt.Println(tabularClusterInfo)

}
