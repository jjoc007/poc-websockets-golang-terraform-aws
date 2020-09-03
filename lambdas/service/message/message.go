package messageservice

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	repositorywebsocket "github.com/jjoc007/poc-websockets-golang-terraform-aws/repository/websocket"
	"os"

	messagemodel "github.com/jjoc007/poc-websockets-golang-terraform-aws/model/message"

	"github.com/jjoc007/poc-websockets-golang-terraform-aws/log"
)

// MessageService describes the structure a message service.
type MessageService interface {
	SendMessage(context.Context, messagemodel.WrapperMessage) error
}

// New creates and returns a new lock service instance
func New(repWs repositorywebsocket.WebSocketRepository) MessageService {
	return &messageService{
		repositoryWebsockets: repWs,
		region:               os.Getenv("AWS_REGION"),
		apiGatewayID:         os.Getenv("API_GATEWAY_ID"),
		environment:          os.Getenv("ENVIRONMENT"),
	}

}

type messageService struct {
	repositoryWebsockets repositorywebsocket.WebSocketRepository
	region               string
	apiGatewayID         string
	environment          string
}

func (s *messageService) SendMessage(ctx context.Context, resource messagemodel.WrapperMessage) (err error) {
	mySession := session.Must(session.NewSession())
	svc := apigatewaymanagementapi.New(mySession, &aws.Config{
		Endpoint: aws.String(fmt.Sprintf("https://%s.execute-api.%s.amazonaws.com/%s", s.apiGatewayID, s.region, s.environment)),
		Region:   aws.String(s.region),
	})

	websocketConnections, err := s.repositoryWebsockets.GetAll(ctx)
	for _, connection := range websocketConnections {
		postToConnectionInput := &apigatewaymanagementapi.PostToConnectionInput{
			ConnectionId: aws.String(connection.ID),
			Data:         []byte(resource.Message),
		}
		_, err = svc.PostToConnection(postToConnectionInput)
		if err != nil {
			log.Logger.Error().Msg(err.Error())
			return
		}
	}
	return
}
