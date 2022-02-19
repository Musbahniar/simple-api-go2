package models

import (
	"context"
	"simple-api-go2/api/utils"
)

type DaftarDiskon struct {
	IdDaftarDiskon int64                  `json:"iddaftardiskon,omitempty"`
	NilaiDiskon    int64                  `json:"nilai,omitempty"`
	TanggalStart   *utils.MysqlFormatDate `json:"tglstart,omitempty"`
	TanggalAkhir   *utils.MysqlFormatDate `json:"tglakhir,omitempty"`
	TahunAjaran    string                 `json:"tahunajaran,omitempty"`
	JenisDiskonRes *JenisDiskonData       `json:"jenis_diskon,omitempty"`
}

type JenisDiskonData struct {
	IdJenisDiskon   int64  `json:"id,omitempty"`
	InisialDiskon   string `json:"inisialdiskon"`
	IdBidang        int64  `json:"idbidang"`
	NamaJenisDiskon string `json:"namajenisdiskon"`
	IsRelatif       string `json:"isrelatif"`
	PartDiskon      string `json:"partdiskon"`
	Prioritas       int64  `json:"prioritas"`
	Kelompk         string `json:"kelompok"`
	Dokumen         string `json:"dokumen"`
	Ikatan          string `json:"ikatan"`
	Status          string `json:"status"`
}

type DaftarDiskonRepo interface {
	GetByTahunAjaran(ctx context.Context, tp string) ([]*DaftarDiskon, error)
	GetOne(ctx context.Context, id string) (*DaftarDiskon, error)
	// GetSearchName(ctx context.Context, namaskl string) ([]*Sekolah, error)
}
