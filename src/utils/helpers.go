package infra

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

// CloudwatchFunc is a function that returns a golang.org/x/net/context.Context
type CloudwatchFunc func() (context.Context, func() error)

// Cloudwatch returns a new Cloudwatch client
func Cloudwatch() (cloudwatch.CloudWatch, func() error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")}, nil)
	if err != nil {
		return nil, func() error { return err }
	}

	return cloudwatch.New(sess), func() error {
		sess.Close()
		return nil
	}
}

// GetMetricData retrieves metric data from AWS Cloudwatch
func GetMetricData(ctx context.Context, client cloudwatch.CloudWatch, metricName string) (*cloudwatch.GetMetricDataOutput, error) {
	params := &cloudwatch.GetMetricDataInput{
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			&cloudwatch.MetricDataQuery{
				Id: aws.String("1"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: cloudwatch.MetricDataQueryMetric{
						Namespace:  aws.String("AWS/RDS"),
						MetricName: aws.String(metricName),
					},
					Period:     aws.Int64(300),
					Stat:       cloudwatch.MetricStatSampleCount,
					Unit:       aws.String("Count"),
				},
				AccountId: aws.String("123456789012"),
				Expression: aws.String("MAX(METRIC_NAME){1,3}"),
				Label: map[string]string{
					"DBInstanceIdentifier": "db-foobar",
				},
				ReturnData: aws.Bool(true),
				Period:     aws.Int64(300),
			},
		},
		Namespace:  aws.String("AWS/RDS"),
		Dimensions: []*cloudwatch.Dimension{},
		StartTime:  aws.Time(time.Now().Add(time.Minute * -15)),
		EndTime:   aws.Time(time.Now()),
	}

	return client.GetMetricDataWithContext(ctx, params)
}

// GetMetricDataWithCallback retrieves metric data from AWS Cloudwatch
func GetMetricDataWithCallback(ctx context.Context, client cloudwatch.CloudWatch, metricName string, callback func(*cloudwatch.GetMetricDataOutput, error)) {
	params := &cloudwatch.GetMetricDataInput{
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			&cloudwatch.MetricDataQuery{
				Id: aws.String("1"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: cloudwatch.MetricDataQueryMetric{
						Namespace:  aws.String("AWS/RDS"),
						MetricName: aws.String(metricName),
					},
					Period:     aws.Int64(300),
					Stat:       cloudwatch.MetricStatSampleCount,
					Unit:       aws.String("Count"),
				},
				AccountId: aws.String("123456789012"),
				Expression: aws.String("MAX(METRIC_NAME){1,3}"),
				Label: map[string]string{
					"DBInstanceIdentifier": "db-foobar",
				},
				ReturnData: aws.Bool(true),
				Period:     aws.Int64(300),
			},
		},
		Namespace:  aws.String("AWS/RDS"),
		Dimensions: []*cloudwatch.Dimension{},
		StartTime:  aws.Time(time.Now().Add(time.Minute * -15)),
		EndTime:   aws.Time(time.Now()),
	}

	_, err := client.GetMetricDataWithContext(ctx, params)
	if err != nil {
		callback(nil, err)
		return
	}

	callback(nil, nil)
}