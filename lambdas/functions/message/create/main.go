package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	messagemodel "github.com/jjoc007/poc-websockets-golang-terraform-aws/model/message"
	servicemessage "github.com/jjoc007/poc-websockets-golang-terraform-aws/service/message"

	"github.com/jjoc007/poc-websockets-golang-terraform-aws/functions"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/log"
	"github.com/jjoc007/poc-websockets-golang-terraform-aws/util"

	"github.com/aws/aws-lambda-go/lambda"
)

func LambdaHandler(cxt context.Context, event events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Logger.Debug().Msg("Start lambda createDynamoTrigger message")
	var messagePayload messagemodel.WrapperMessage
	body := event.Body
	err := json.Unmarshal([]byte(body), &messagePayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on decoding body %v", messagePayload)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}

	err = functions.Instances["messageService"].(servicemessage.MessageService).SendMessage(cxt, messagePayload)
	if err != nil {
		log.Logger.Error().Err(err).Msgf("ERROR on the createDynamoTrigger message %v", messagePayload)
		return util.ResponseErrorFunction(err, fmt.Sprintf("Error when it is process request")), err
	}
	return events.APIGatewayProxyResponse{StatusCode: 200}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
