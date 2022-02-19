package models

import (
	"context"
	"simple-api-go2/api/utils"
)

type TahunAjaran struct {
	TahunAjaran string                 `json:"tahunajaran,omitempty"`
	IsDefault   string                 `json:"isdefault,omitempty"`
	Awal        *utils.MysqlFormatDate `json:"tglawal,omitempty"`
	Akhir       *utils.MysqlFormatDate `json:"tglakhir"`
}

type TahunAjaranRepo interface {
	GetAll(ctx context.Context, limit int64, offset int64) ([]*TahunAjaran, error)
}
