package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/xid"
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

type EC2DescribeInstancesAPI interface {
	DescribeInstances(ctx context.Context,
		params *ec2.DescribeInstancesInput,
		optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

// GetInstances retrieves information about your Amazon Elastic Compute Cloud (Amazon EC2) instances.
// Inputs:
//     c is the context of the method call, which includes the AWS Region.
//     api is the interface that defines the method call.
//     input defines the input arguments to the service call.
// Output:
//     If success, a DescribeInstancesOutput object containing the result of the service call and nil.
//     Otherwise, nil and an error from the call to DescribeInstances.
func GetInstances(c context.Context, api EC2DescribeInstancesAPI, input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	return api.DescribeInstances(c, input)
}

func DescribeInstancesCmd() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ec2.NewFromConfig(cfg)

	input := &ec2.DescribeInstancesInput{}

	result, err := GetInstances(context.TODO(), client, input)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err)
		return
	}

	for _, r := range result.Reservations {
		fmt.Println("Reservation ID: " + *r.ReservationId)
		fmt.Println("Instance IDs:")
		for _, i := range r.Instances {
			fmt.Println("   " + *i.InstanceId)
		}

		fmt.Println("")
	}
}

func main() {

	myBucketName := "mybucket-" + (xid.New().String())

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		panic("Failed to load configuration")
	}

	s3client := s3.NewFromConfig(cfg)

	AccountBucketOps(*s3client, myBucketName)
	DescribeInstancesCmd()
}
