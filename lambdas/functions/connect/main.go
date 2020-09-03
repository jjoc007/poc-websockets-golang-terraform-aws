package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	servicewebsocket "github.com/jjoc007/poc-websockets-golang-terraform-aws/service/websocket"

	"github.com/jjoc007/poc-websockets-golang-terraform-aws/functions"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/log"
	websocketmodel "github.com/jjoc007/poc-websockets-golang-terraform-aws/model/websocket"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/util"

	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda createDynamoTrigger websocket connection")
	log.Logger.Debug().Msgf("Start lambda createDynamoTrigger claims %v", event.RequestContext.Authorizer)
	websocket := &websocketmodel.WebSocket{
		ID:  event.RequestContext.ConnectionID,
	}
	log.Logger.Debug().Msgf("id %v", websocket)
	err := functions.Instances["websocketService"].(servicewebsocket.WebsocketService).Create(cxt, websocket)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the connect %v", websocket)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
