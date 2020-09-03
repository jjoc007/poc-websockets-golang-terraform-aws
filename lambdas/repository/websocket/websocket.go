package repositorywebsocket

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	websocketmodel "github.com/jjoc007/poc-websockets-golang-terraform-aws/model/websocket"

	"github.com/jjoc007/poc-websockets-golang-terraform-aws/config"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/log"
)

// WebSocketRepository describes the lock repository.
type WebSocketRepository interface {
	Create(context.Context, *websocketmodel.WebSocket) error
	GetAll(context.Context) ([]websocketmodel.WebSocket, error)
	Delete(context.Context, string) error
}

// NewRepository creates and returns a new Websocket repository instance
func NewRepository(database *dynamodb.DynamoDB) WebSocketRepository {
	return &repository{
		database: database,
		table:    config.POCWebsocketConnectionTable,
	}
}

type repository struct {
	database *dynamodb.DynamoDB
	table    string
}

func (s *repository) Create(ctx context.Context, resource *websocketmodel.WebSocket) (err error) {
	log.Logger.Debug().Msgf("Adding a new notification [%s] ", resource.ID)

	av, err := dynamodbattribute.MarshalMap(resource)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = s.database.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(s.table),
		Item:      av,
	})

	log.Logger.Debug().Msgf("ID %v inserted.\n", resource.ID)
	return nil
}

func (s *repository) GetAll(ctx context.Context) (connections []websocketmodel.WebSocket, err error) {
	log.Logger.Debug().Msgf("Getting connections")
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, config.Timeout)
		defer cancel()
	}

	input := &dynamodb.ScanInput{
		TableName: aws.String(s.table),
	}

	result, err := s.database.Scan(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, item := range result.Items {
		websocket := &websocketmodel.WebSocket{}
		err = dynamodbattribute.UnmarshalMap(item, &websocket)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		connections = append(connections, *websocket)
	}

	return connections, nil
}

func (s repository) Delete(ctx context.Context, id string) (err error) {
	log.Logger.Debug().Msgf("Deleting a connection [%s] ", id)
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, config.Timeout)
		defer cancel()
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"connectionId": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(s.table),
	}

	_, err = s.database.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	return nil
}
