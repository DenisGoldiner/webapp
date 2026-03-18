package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/DenisGoldiner/webapp/internal"
)

type Client struct {
	dbExec sqlx.ExtContext
}

func NewClient(dbExec sqlx.ExtContext) Client {
	return Client{dbExec: dbExec}
}

func (c Client) GetTraveller(ctx context.Context, id uuid.UUID) (internal.Traveller, error) {
	q := "select id, first_name, last_name, age from travellers where id = $1"

	rows, err := c.dbExec.QueryxContext(ctx, q, id)
	if err != nil {
		return internal.Traveller{}, fmt.Errorf("failed to fetch traveler: %w", err)
	}

	defer func() { _ = rows.Close() }()

	var travelers []Traveller

	for rows.Next() {
		var traveler Traveller

		if err = rows.StructScan(&traveler); err != nil {
			return internal.Traveller{}, fmt.Errorf("failed to scan traveler: %w", err)
		}

		travelers = append(travelers, traveler)
	}

	if len(travelers) == 0 {
		return internal.Traveller{}, fmt.Errorf("no travelers with id %s: %w", id, internal.ErrNoResource)
	}

	return internal.Traveller{
		ID:        travelers[0].ID,
		FirstName: travelers[0].FirstName,
		LastName:  travelers[0].LastName,
		Age:       travelers[0].Age,
	}, err
}

func (c Client) CreateTraveller() {

}

func (c Client) DeleteTraveller() {

}
