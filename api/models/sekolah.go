package models

import (
	"context"
)

type Sekolah struct {
	IdSekolah   int64    `json:"idsekolah,omitempty"`
	NamaSekolah string   `json:"namasekolah,omitempty"`
	Jenjang     string   `json:"jenjang,omitempty"`
	ProvinsiRes *GeoData `json:"provinsi_details,omitempty"`
}

type GeoData struct {
	IdProvinsi   *int64  `json:"idprovinsi,omitempty"`
	NamaProvinsi *string `json:"namaprovinsi,omitempty"`
	IdDistrict   *int64  `json:"iddistrict,omitempty"`
	NamaDistrict *string `json:"namadistrict,omitempty"`
}

type SekolahRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*Sekolah, error)
	GetOne(ctx context.Context, id string) (*Sekolah, error)
	GetSearchName(ctx context.Context, namaskl string) ([]*Sekolah, error)
}
