package internal

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrNoResource = errors.New("no resource found")

type TravellerStorage interface {
	Get(ctx context.Context, id uuid.UUID) (Traveller, error)
	Create(ctx context.Context, params CreateTravellerPayload) (uuid.UUID, error)
}

type Traveller struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Age       int
}

type CreateTravellerPayload struct {
	FirstName string
	LastName  string
	Age       int
}
