package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"gopkg.in/alecthomas/kingpin.v2"
	"strings"
	"time"
)

var (
	region     = kingpin.Flag("region", "The region to use.").Required().String()
	namespace  = kingpin.Flag("namespace", "The namespace to use").Required().String()
	dimensions = kingpin.Flag("dimensions", "The dimensions to use").Default("").String()
	metric     = kingpin.Flag("metric", "The metric to use").Required().String()
	unit       = kingpin.Flag("unit", "The unit to use").Required().String()
	value      = kingpin.Flag("value", "The value to use").Required().Float64()
	resolution = kingpin.Flag("resolution", "The resolution to use").Default("60").Int64()
)

func main() {
	kingpin.Parse()

	// Create the AWS session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(*region),
	}))

	dims := make([]*cloudwatch.Dimension, 0, 0)
	for _, arg := range strings.Split(*dimensions, ",") {
		parts := strings.Split(arg, "=")
		d := &cloudwatch.Dimension{}
		d.SetName(parts[0])
		d.SetValue(parts[1])
		dims = append(dims, d)
	}

	now := time.Now()
	datum := &cloudwatch.MetricDatum{
		MetricName:        aws.String(*metric),
		Unit:              aws.String(*unit),
		Dimensions:        dims,
		Timestamp:         &now,
		Value:             value,
		StorageResolution: resolution,
	}

	datums := []*cloudwatch.MetricDatum{datum}

	// Create a cloudwatch client
	cw := cloudwatch.New(sess)

	// Upload the file to cloud watch
	_, err := cw.PutMetricData(&cloudwatch.PutMetricDataInput{
		MetricData: datums,
		Namespace:  aws.String(*namespace),
	})

	// Check for errors
	if err != nil {
		fmt.Printf("Failed to publish metric, %v\n", err)
		return
	}
}
