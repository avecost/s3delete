package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"os"
)

func main() {

	c, err := GetConfig("./conf/audiofp.conf")
	if err != nil {
		fmt.Println("Config file not found")
		return
	}

	//bucket := os.Args[1]
	fileToRemove := os.Args[2]

	// Initialize a client using Spaces
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(c.Bucket.AccessKey, c.Bucket.SecretKey, ""),
		Endpoint:    aws.String(c.Bucket.Endpoint),
		Region:      aws.String(c.Bucket.Region),
	}

	newSession := session.New(s3Config)
	svc := s3.New(newSession)
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(c.Bucket.Name),
		Key:    aws.String(fileToRemove),
	}

	_, err = svc.DeleteObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}

	fmt.Printf("Successfully remove %s\n", fileToRemove)
}
