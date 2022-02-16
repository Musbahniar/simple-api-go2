package models

import (
	"context"
)

type SekolahKelas struct {
	IdSekolahKelas   int64  `json:"idsekolahkelas,omitempty"`
	NamaSekolahKelas string `json:"namasekolahkelas,omitempty"`
	TingkatKelas     int64  `json:"tingkatkelas,omitempty"`
	Jurusan          string `json:"jurusankelas,omitempty"`
	KelompokKelas    string `json:"kelompokkelas,omitempty"`
}

type SekolahKelasRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*SekolahKelas, error)
	GetOne(ctx context.Context, id string) (*SekolahKelas, error)
	GetByKelompok(ctx context.Context, kelompok string) ([]*SekolahKelas, error)
}
