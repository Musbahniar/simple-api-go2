package controllers

import (
	"simple-api-go2/api/models"
	r "simple-api-go2/api/repositories"
	"simple-api-go2/config/driver"
)

type TahunAjaran struct {
	tahunajaranRepo models.TahunAjaran
}

func NewTahunAjaranHandler(db *driver.DB) *TahunAjaran {
	return &TahunAjaran{
		tahunajaranRepo: r.NewTahunAjaranRepo(db.SQL),
	}
}
