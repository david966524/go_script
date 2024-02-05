package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func main() {
	key := ""
	secret := ""
	region := "ap-east-1"
	config := aws.Config{
		Credentials: credentials.NewStaticCredentialsProvider(key, secret, ""),
		Region:      region,
	}
	//get
	getEc2Id(config)
	//deleteEc2()
	//createEc2()
}

func getEc2Id(config aws.Config) []string {
	client := ec2.NewFromConfig(config)
	instances, err := client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		fmt.Println(err.Error())
	}
	var ec2Ids []string
	// 遍历返回id
	for _, reservation := range instances.Reservations {
		for _, instance := range reservation.Instances {
			ec2Ids = append(ec2Ids, *instance.InstanceId)
		}
	}
	return ec2Ids
}
func deleteEc2(config aws.Config, ec2ids []string) {
	client := ec2.NewFromConfig(config)
	result, err := client.TerminateInstances(context.TODO(), &ec2.TerminateInstancesInput{
		InstanceIds: ec2ids,
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, i := range result.TerminatingInstances {
		fmt.Println("Success delete ", *i.InstanceId)
	}

}
func createEc2(config aws.Config) {
	DEVICE_NAME := "/dev/xvda"
	var DISK_SIZE_GB int32 = 10
	imageId := "ami-0f4dfdbd0b33c2e4c"
	instenceType := "t3.small"
	keyName := "aws_key"
	subnetId := "subnet-0de0877c406328034"
	securityGroupIds := []string{"sg-0fc4badd3cbf266f7"}
	TAG_NAME := "Name"
	TAG_VALUE := "myapitest6"
	_ = TAG_VALUE
	//arn := "arn:aws:iam::293116843282:role/system_manager"
	rolename := "system_manager"
	role := types.IamInstanceProfileSpecification{
		//Arn:  &arn,    //arn 和 name 不能同时使用
		Name: &rolename,
	}
	tag := []types.TagSpecification{
		{
			ResourceType: "instance",
			Tags: []types.Tag{
				types.Tag{
					Key:   &TAG_NAME,
					Value: &TAG_VALUE,
				},
			},
		},
	}
	blockDeviceMappings := []types.BlockDeviceMapping{
		{
			DeviceName: &DEVICE_NAME,
			Ebs: &types.EbsBlockDevice{
				//Encrypted:           nil,
				VolumeSize: &DISK_SIZE_GB,
				VolumeType: "gp3",
			},
		},
	}
	userdata := fmt.Sprintf(`#! /bin/bash
yum update -y
yum install httpd -y
echo 'userdata script' > /test.txt
hostnamectl set-hostname --static %v`, TAG_VALUE)
	baseUserData := base64.StdEncoding.EncodeToString([]byte(userdata))
	var count int32 = 1
	//dryrun := true
	clinet := ec2.NewFromConfig(config)
	instances, err := clinet.RunInstances(context.TODO(), &ec2.RunInstancesInput{
		MaxCount:            &count,
		MinCount:            &count,
		BlockDeviceMappings: blockDeviceMappings,
		IamInstanceProfile:  &role,
		//DryRun:              &dryrun,
		ImageId:           &imageId,
		InstanceType:      types.InstanceType(instenceType),
		KeyName:           &keyName,
		SecurityGroupIds:  securityGroupIds,
		SubnetId:          &subnetId,
		TagSpecifications: tag,
		UserData:          &baseUserData,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	for _, i := range instances.Instances {
		fmt.Println(fmt.Sprintf("ec2_ID：%v", *i.InstanceId))
	}
}
