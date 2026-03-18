package internal

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type Travellers struct {
	db Storage
}

func NewTravellers(db Storage) Travellers {
	return Travellers{db: db}
}

func (t Travellers) GetTraveller(ctx context.Context, id uuid.UUID) (Traveller, error) {
	if id == uuid.Nil {
		return Traveller{}, fmt.Errorf("id must be a valid uuid")
	}

	res, err := t.db.GetTraveller(ctx, id)
	if err != nil {
		return Traveller{}, fmt.Errorf("%w: failed to get traveller from db", err)
	}

	return res, nil
}

func (t Travellers) CreateTraveller(ctx context.Context, traveller Traveller) (uuid.UUID, error) {
	slog.Info("create traveler payload", "traveler", traveller)
	return uuid.New(), nil
}

func (t Travellers) DeleteTraveller() {

}
