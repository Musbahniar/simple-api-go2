package models

import (
	"context"
	"database/sql"
)

type TahunAjaran struct {
	TahunAjaran string       `json:"tahunajaran,omitempty"`
	IsDefault   string       `json:"isdefault,omitempty"`
	Awal        sql.NullTime `json:"tglawal,omitempty"`
	Akhir       sql.NullTime `json:"tglakhir"`
}

type TahunAjaranRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*TahunAjaran, error)
}
