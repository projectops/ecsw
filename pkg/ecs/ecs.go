package ecs

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

// Services - services from ecs
type Services struct {
	ARN          string
	Name         string
	DesiredTasks int64
	RunningTasks int64

	Tasks Tasks
}

// Tasks - tasks from services
type Tasks struct {
	ARN    []string
	Length int64
}

func auth(region string) (*ecs.ECS, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	svc := ecs.New(sess)

	return svc, nil
}

func listServices(cluster, region string) (*ecs.ListServicesOutput, error) {
	svc, err := auth(region)
	if err != nil {
		log.Printf("Auth Problem: %s\n", err)
		return nil, err
	}

	input := &ecs.ListServicesInput{
		Cluster:    aws.String(cluster),
		MaxResults: aws.Int64(100),
	}

	result, err := svc.ListServices(input)

	if err != nil {
		return nil, err
	}

	return result, err
}

func listTasks(cluster, service, region string) (*ecs.ListTasksOutput, error) {
	svc, err := auth(region)
	if err != nil {
		log.Printf("Auth Problem: %s\n", err)
		return nil, err
	}

	input := &ecs.ListTasksInput{
		Cluster:     aws.String(cluster),
		ServiceName: aws.String(service),
	}

	result, err := svc.ListTasks(input)
	if err != nil {
		log.Printf("Listing task problem: %s\n", err)
		return nil, err

	}

	return result, nil
}

func descService(cluster, arn, region string) (*ecs.DescribeServicesOutput, error) {
	svc, err := auth(region)
	if err != nil {
		log.Printf("Auth Problem: %s\n", err)
		return nil, err
	}

	input := &ecs.DescribeServicesInput{
		Cluster: aws.String(cluster),
		Services: []*string{
			aws.String(arn),
		},
	}

	result, err := svc.DescribeServices(input)
	if err != nil {
		log.Printf("Describe service problem: %s\n", err)
		return nil, err
	}

	return result, nil
}
