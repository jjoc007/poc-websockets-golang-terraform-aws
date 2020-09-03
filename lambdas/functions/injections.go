package functions

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/config/db"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/log"
	repositorywebsocket "github.com/jjoc007/poc-websockets-golang-terraform-aws/repository/websocket"
	servicemessage "github.com/jjoc007/poc-websockets-golang-terraform-aws/service/message"
	servicewebsocket "github.com/jjoc007/poc-websockets-golang-terraform-aws/service/websocket"
)

// Instances is a global map that contain all object instances of app
var Instances = MakeDependencyInjection()

// MakeDependencyInjection Initialize all dependencies
func MakeDependencyInjection() map[string]interface{} {
	log.Logger.Debug().Msg("Start bootstrap app objects")
	instances := make(map[string]interface{})

	database, err := db.NewDynamoDBStorage()
	if err != nil {
		panic(err)
	}
	instances["dataBase"] = database

	instances["websocketRepository"] = repositorywebsocket.NewRepository(database.GetConnection().(*dynamodb.DynamoDB))

	instances["websocketService"] = servicewebsocket.New(
		instances["websocketRepository"].(repositorywebsocket.WebSocketRepository))

	instances["messageService"] = servicemessage.New(
		instances["websocketRepository"].(repositorywebsocket.WebSocketRepository))

	log.Logger.Debug().Msg("End bootstrap app objects")
	return instances
}
