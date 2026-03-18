package internal

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrNoResource = errors.New("no resource found")

type Storage interface {
	GetTraveller(ctx context.Context, id uuid.UUID) (Traveller, error)
}

type Traveller struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Age       int
}
