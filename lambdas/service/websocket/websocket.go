package servicewebsocket

import (
	"context"

	websocketmodel "github.com/jjoc007/poc-websockets-golang-terraform-aws/model/websocket"

	"github.com/jjoc007/poc-websockets-golang-terraform-aws/log"
	repositorywebsocket "github.com/jjoc007/poc-websockets-golang-terraform-aws/repository/websocket"

	"github.com/pkg/errors"
)

// WebsocketService describes the structure a websocket service.
type WebsocketService interface {
	Create(context.Context, *websocketmodel.WebSocket) error
	Delete(context.Context, string) error
}

// New creates and returns a new lock service instance
func New(rep repositorywebsocket.WebSocketRepository) WebsocketService {
	return &websocketService{
		repository: rep,
	}
}

type websocketService struct {
	repository repositorywebsocket.WebSocketRepository
}

func (s *websocketService) Create(ctx context.Context, resource *websocketmodel.WebSocket) error {
	log.Logger.Debug().Msg("Creating web socket connection")
	err := s.repository.Create(ctx, resource)
	if err != nil {
		return errors.Wrapf(err, "Error creating websocket connection on services [%s]", resource.ID)
	}
	return nil
}

func (s *websocketService) Delete(ctx context.Context, id string) error {
	log.Logger.Debug().Msg("Deleting a websocket connection on services")
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return errors.Wrapf(err, "Deleting a websocket connection on services [%s]", id)
	}
	return nil
}
