package main

import (
	"fmt"
	"time"

	"github.com/luck02/goPublish/vendor/github.com/aws/aws-sdk-go/aws"
	"github.com/luck02/goPublish/vendor/github.com/aws/aws-sdk-go/service/kinesis"
)

func main() {
	svc := kinesis.New(&aws.Config{Region: aws.String("us-east-1")})

	testStream := "testStream"
	CreateStream(svc, testStream)
	CheckStreamStatus("ACTIVE", svc, testStream)
	PublishMessage()
	DeleteStream(svc, testStream)

	descParams := &kinesis.DescribeStreamInput{
		StreamName: aws.String(testStream),
	}

	respDescribe, err := svc.DescribeStream(descParams)

	fmt.Println(respDescribe)
	fmt.Println(err)

}

func CheckStreamStatus(status string, service *kinesis.Kinesis, streamName string) bool {

	descParams := &kinesis.DescribeStreamInput{
		StreamName: aws.String(streamName),
	}

	for i := 0; i < 30; i++ {
		respDescribe, err := service.DescribeStream(descParams)

		if *respDescribe.StreamDescription.StreamStatus == status {
			return true
		}
		fmt.Println(respDescribe)
		fmt.Println(err)

		time.Sleep(500 * time.Millisecond)

	}

	return false
}

func CreateStream(service *kinesis.Kinesis, streamName string) {
	params := &kinesis.CreateStreamInput{
		ShardCount: aws.Int64(1),
		StreamName: aws.String(streamName),
	}

	resp, err := service.CreateStream(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Response from CreateStream:")
	fmt.Println(resp)
	fmt.Println(err)
}

func PublishMessage() {
	svc := kinesis.New(nil)
	fmt.Println("derp:")
	fmt.Println(svc)

	fmt.Println("this is the der")
	fmt.Println("derp")

}

func DeleteStream(service *kinesis.Kinesis, streamName string) {
	params := &kinesis.DeleteStreamInput{
		StreamName: aws.String(streamName),
	}

	resp, err := service.DeleteStream(params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Response from deleteStream")
	fmt.Println(resp)

}
