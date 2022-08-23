package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/ryanuber/columnize"
)

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

func DescribeInstancesCmd(key string) {
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
	var Output []string
	var tableFormat = "Instance ID | Flavour | KeyName | Age | State| Tag-Key | Tag-Value"
	Output = append(Output, tableFormat)
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			var tag, value, keyName string = "-", "-", "-"
			age := time.Since(*i.LaunchTime).Round(time.Second).String()
			if i.Tags != nil {
				tag = *i.Tags[0].Key
				value = *i.Tags[0].Value
			}
			if i.KeyName != nil {
				keyName = *i.KeyName
			}
			if key == "" {
				instanceInfo := fmt.Sprintf("%v | %v | %v | %v | %v | %v | %v", *i.InstanceId, i.InstanceType, keyName, age, i.State.Name, tag, value)
				Output = append(Output, instanceInfo)
			} else {
				if key == keyName {
					instanceInfo := fmt.Sprintf("%v | %v | %v | %v | %v | %v | %v", *i.InstanceId, i.InstanceType, keyName, age, i.State.Name, tag, value)
					Output = append(Output, instanceInfo)
				}
			}

		}
	}
	tabularClusterInfo := columnize.SimpleFormat(Output)
	fmt.Println(tabularClusterInfo)
}
