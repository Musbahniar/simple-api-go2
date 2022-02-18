package models

import "context"

type CreateJenisDiskon struct {
	InisialDiskon   string `json:"inisialdiskon" validate:"required,min=2,max=3"`
	IdBidang        int64  `json:"idbidang" validate:"required,number"`
	NamaJenisDiskon string `json:"namajenisdiskon" validate:"required"`
	IsRelatif       string `json:"isrelatif" validate:"required,oneof=Y N"`
	PartDiskon      string `json:"partdiskon" validate:"required,oneof=Y N"`
	Prioritas       int64  `json:"prioritas" validate:"required,gte=1"`
	Kelompk         string `json:"kelompok" validate:"required"`
	Dokumen         string `json:"dokumen" validate:"required,oneof=Y N"`
	Ikatan          string `json:"ikatan" validate:"required,oneof=SISWA CABANG PUSAT"`
	Status          string `json:"status" validate:"required,oneof=Aktif Create Submit"`
}

type CreateJenisDiskonResponse struct {
	IdJenisDiskon   int64  `json:"id,omitempty"`
	InisialDiskon   string `json:"inisialdiskon,omitempty"`
	NamaJenisDiskon string `json:"namajenisdiskon,omitempty"`
}

type JenisDiskon struct {
	IdJenisDiskon   int64  `json:"id,omitempty" validate:"required"`
	InisialDiskon   string `json:"inisialdiskon" validate:"required,min=2,max=3"`
	IdBidang        int64  `json:"idbidang"`
	NamaJenisDiskon string `json:"namajenisdiskon" validate:"required"`
	IsRelatif       string `json:"isrelatif" validate:"required,oneof=Y N"`
	PartDiskon      string `json:"partdiskon" validate:"required,oneof=Y N"`
	Prioritas       int64  `json:"prioritas" validate:"required,gte=1"`
	Kelompk         string `json:"kelompok" validate:"required"`
	Dokumen         string `json:"dokumen" validate:"required,oneof=Y N"`
	Ikatan          string `json:"ikatan" validate:"required,oneof=SISWA CABANG PUSAT"`
	Status          string `json:"status" validate:"required,oneof=Aktif Create Submit"`
}

type JenisDiskonRepo interface {
	Create(ctx context.Context, g *CreateJenisDiskon) (*CreateJenisDiskonResponse, error)
	Update(ctx context.Context, g *JenisDiskon) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetAll(ctx context.Context, limit int64, offset int64) ([]*JenisDiskon, error)
	GetOne(ctx context.Context, id string) (*JenisDiskon, error)
}
