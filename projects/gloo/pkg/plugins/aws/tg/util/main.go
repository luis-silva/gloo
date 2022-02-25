package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/tg"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
)

var accessKey = flag.String("accesskey", "", "access key id")
var secretKey = flag.String("secretkey", "", "secret key")
var region = flag.String("region", "", "secret key")
var debug = flag.Bool("debug", false, "debug")

func main() {
	flag.Parse()
	config := &aws.Config{Region: aws.String(*region)}
	if *debug {
		config = config.WithLogLevel(aws.LogDebugWithHTTPBody)
	}
	sess, err := session.NewSession(config.
		WithCredentials(credentials.NewStaticCredentials(*accessKey, *secretKey, "")))

	if err != nil {
		panic(err)
	}
	elbv2api := awselbv2.New(sess)

	e := tg.NewTargetGroupLister()

	targetgroups, err := e.ListWithClient(context.Background(), elbv2api)
	if err != nil {
		panic(err)
	}

	fmt.Printf("found %d targetgroups \n", len(targetgroups))

	tagdescriptions, err := e.ListTagsWithClient(targetgroups, context.Background(), elbv2api)
	if err != nil {
		panic(err)
	}

	fmt.Printf("found %d tagdescriptions \n", len(tagdescriptions))

	tags := e.ConvertTagDescriptionsToMap(tagdescriptions)

	extendertargetgroups := e.MergeTargetGroupTags(targetgroups, tags)

	for _, tg := range extendertargetgroups {
		for _, tag := range tg.Tags {
			if *tag.Key == "Name" {
				fmt.Printf("%v\n", *tag.Value)
			}
		}
	}
}
