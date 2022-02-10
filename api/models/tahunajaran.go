package models

import (
	"context"
	"database/sql"
)

type TahunAjaran struct {
	TahunAjaran string       `json:"c_TahunAjaran,omitempty"`
	IsDefault   string       `json:"c_IsDefault,omitempty"`
	Awal        sql.NullTime `json:"c_awal,omitempty"`
	Akhir       sql.NullTime `json:"c_akhir"`
}

type TahunAjaranRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*TahunAjaran, error)
	// GetOne(ctx context.Context, id string) (*Bidang, error)
}
