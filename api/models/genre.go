package models

import (
	"context"
	"time"
)

type Genre struct {
	Id        *int64     `json:"id,string,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Slug      *string    `json:"slug,omitempty"`
	Status    *bool      `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type GenreRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*Genre, error)
	GetOne(ctx context.Context, id string) (*Genre, error)
}
