package repository

import (
	"context"
	"echo-template/db"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrClientNotFound = errors.New("client not found")

type ClientRepository struct {
	db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *ClientRepository {
	if db == nil {
		panic("Database connection is nil in repository")
	}
	return &ClientRepository{db: db}
}

func (r *ClientRepository) CreateClient(ctx context.Context, params db.CreateClientParams) (*db.Client, error) {
	q := db.New(r.db)
	client, err := q.CreateClient(ctx, params)
	return &client, err
}

func (r *ClientRepository) GetClientByEmail(ctx context.Context, phoneNumber string) (*db.Client, error) {
	q := db.New(r.db)
	client, err := q.GetClient(ctx, phoneNumber)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrClientNotFound
	}
	return &client, err
}
