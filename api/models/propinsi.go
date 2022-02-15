package models

import (
	"context"
)

type Provinsi struct {
	IdProvinsi   int64  `json:"idprovinsi,omitempty"`
	NamaProvinsi string `json:"namaprovinsi,omitempty"`
}

type ProvinsiRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*Provinsi, error)
	GetOne(ctx context.Context, id string) (*Provinsi, error)
}
