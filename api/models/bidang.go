package models

import (
	"context"
	"database/sql"
)

type Bidang struct {
	IdBidang      int64         `json:"c_IdBidang,omitempty"`
	NamaBidang    string        `json:"c_NamaBidang,omitempty"`
	Status        string        `json:"c_Status,omitempty"`
	IdKewilayahan int64         `json:"c_IdKewilayahan,omitempty"`
	Upline        sql.NullInt64 `json:"c_Upline"`
}

type BidangRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*Bidang, error)
	GetOne(ctx context.Context, id string) (*Bidang, error)
}
