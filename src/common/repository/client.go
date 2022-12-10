package repository

import (
	"context"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var a sync.Once
var client *dynamodb.Client

func GetClient() (dc *dynamodb.Client, err error) {
	a.Do(func() {
		c, err := config.LoadDefaultConfig(
			context.Background(),
			config.WithRegion(getRegion()),
		)

		if err != nil {
			return
		}

		dc = dynamodb.NewFromConfig(c)
		client = dc
	})

	return
}

func getRegion() string {
	if env := os.Getenv("REGION"); env != "" {
		return env
	}

	return "us-east-1"
}
