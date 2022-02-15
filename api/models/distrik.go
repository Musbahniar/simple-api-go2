package models

import (
	"context"
)

type Distrik struct {
	IdDistrik   int64         `json:"iddistrik,omitempty"`
	NamaDistrik string        `json:"namadistrik,omitempty"`
	PhoneCode   string        `json:"phonecode,omitempty"`
	ProvinsiRes *ProvinsiData `json:"provinsi_details,omitempty"`
}

type ProvinsiData struct {
	IdProvinsi   *int64  `json:"idprovinsi,omitempty"`
	NamaProvinsi *string `json:"namaprovinsi,omitempty"`
}

type DistrikRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*Distrik, error)
	GetOne(ctx context.Context, id string) (*Distrik, error)
	GetByIdProvinsi(ctx context.Context, id string) ([]*Distrik, error)
}
