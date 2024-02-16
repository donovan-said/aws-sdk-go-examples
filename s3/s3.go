/*
Examples found in:
- https://aws.github.io/aws-sdk-go-v2/docs/getting-started/
*/

package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func create_s3_bucket() {
    // Load the shared AWS Configuration (~/.aws/config)

    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    // Create S3 service client
    client := s3.NewFromConfig(cfg)

    // Get the first page of results from ListObjectV2 for an S3 bucket
    output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
        Bucket: aws.String("MySampleBucket"),
    })
    if err != nil {
        log.Fatal()
    }

    log.Println("Results from first page: ")
    for _, object := range output.Contents {
        log.Printf("key=%s, size=%d", aws.ToString(object.Key), *object.Size)
    }

}

func main() {
    // Init function
    create_s3_bucket()
}
