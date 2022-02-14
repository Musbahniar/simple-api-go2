package models

import (
	"context"
)

type Penanda struct {
	IdPenanda   int64  `json:"c_IdPenanda,omitempty"`
	NamaPenanda string `json:"c_Penanda,omitempty"`
}

type PenandaRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*Penanda, error)
	GetOne(ctx context.Context, id string) (*Penanda, error)
}
