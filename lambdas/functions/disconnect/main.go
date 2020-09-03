package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	servicewebsocket "github.com/jjoc007/poc-websockets-golang-terraform-aws/service/websocket"

	"github.com/jjoc007/poc-websockets-golang-terraform-aws/functions"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/log"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/util"

	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	id := event.RequestContext.ConnectionID
	log.Logger.Debug().Msg("Start lambda delete websocket connection")
	log.Logger.Debug().Msgf("id %v", id)
	err := functions.Instances["websocketService"].(servicewebsocket.WebsocketService).Delete(cxt, id)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the connect %v", id)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
