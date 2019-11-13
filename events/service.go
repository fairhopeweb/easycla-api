package events

import (
	"context"

	"github.com/communitybridge/easycla-api/gen/models"
	"github.com/communitybridge/easycla-api/gen/restapi/operations/events"
)

// Service interface defines methods of event service
type Service interface {
	CreateEvent(event models.Event) error
	ListEvents(ctx context.Context, params *events.ListEventsParams) (*models.EventList, error)
}

type service struct {
	repo Repository
}

// NewService creates new instance of event service
func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateEvent(event models.Event) error {
	return s.repo.CreateEvent(&event)
}

func (s *service) ListEvents(ctx context.Context, params *events.ListEventsParams) (*models.EventList, error) {
	return s.repo.ListEvents(ctx, params)
}
