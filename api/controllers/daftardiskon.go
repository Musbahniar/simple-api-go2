package controllers

import (
	"net/http"
	"simple-api-go2/api/models"
	r "simple-api-go2/api/repositories"
	"simple-api-go2/config/driver"
	"simple-api-go2/handler"

	"github.com/gorilla/mux"
)

type DaftarDiskon struct {
	daftarDiskonRepo models.DaftarDiskonRepo
}

func NewDaftarDiskonHandler(db *driver.DB) *DaftarDiskon {
	return &DaftarDiskon{
		daftarDiskonRepo: r.NewDaftarDiskonRepo(db.SQL),
	}
}

func (g *DaftarDiskon) GetByTahunAjaran(w http.ResponseWriter, r *http.Request) {
	//get the query params
	var tp string

	if r.URL.Query().Get("tp") == "" {
		tp = ""
	} else {
		tp = r.URL.Query().Get("tp")
	}

	res, err := g.daftarDiskonRepo.GetByTahunAjaran(r.Context(), tp)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *DaftarDiskon) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, err := g.daftarDiskonRepo.GetOne(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}
