package main

import (
	"fmt"

	"github.com/luck02/goPublish/vendor/github.com/aws/aws-sdk-go/service/kinesis"
)

func main() {
	fmt.Println("I ran")

	PublishMessage()
}

func PublishMessage() {
	svc := kinesis.New(nil)
	fmt.Println("derp:")
	fmt.Println(svc)

}
