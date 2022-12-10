package event

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jeanmolossi/special-broccoli/common/repository"
)

type (
	EventDispatcher struct {
		client *dynamodb.Client
	}
)

func (e *EventDispatcher) Emit(event *Event) error {
	output, err := e.client.PutItem(
		context.Background(),
		&dynamodb.PutItemInput{
			TableName: aws.String("Events"),
			Item:      eventDataToItem(event),
		},
	)

	if err != nil {
		return err
	}

	log.Printf("%+v\n", output)

	return nil
}

func NewEventDispatcher() *EventDispatcher {
	client, err := repository.GetClient()
	if err != nil {
		panic("fail getting dynamo client " + err.Error())
	}

	return &EventDispatcher{
		client: client,
	}
}
