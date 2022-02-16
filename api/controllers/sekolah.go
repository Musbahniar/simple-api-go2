package controllers

import (
	"net/http"
	hc "simple-api-go2/api/constants"
	"simple-api-go2/api/models"
	r "simple-api-go2/api/repositories"
	"simple-api-go2/config/driver"
	"simple-api-go2/handler"
	"strconv"

	"github.com/gorilla/mux"
)

type Sekolah struct {
	sekolahRepo models.SekolahRepo
}

func NewSekolahHandler(db *driver.DB) *Sekolah {
	return &Sekolah{
		sekolahRepo: r.NewSekolahRepo(db.SQL),
	}
}

func (g *Sekolah) GetAll(w http.ResponseWriter, r *http.Request) {
	//get the query params
	var limit, offset int64

	if r.URL.Query().Get("limit") == "" {
		limit = hc.LIMIT
	} else {
		limit, _ = strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	}

	if r.URL.Query().Get("offset") == "" {
		offset = hc.OFFSET
	} else {
		offset, _ = strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)
	}

	res, err := g.sekolahRepo.GetAll(r.Context(), limit, offset)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *Sekolah) GetOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := g.sekolahRepo.GetOne(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *Sekolah) GetSearchName(w http.ResponseWriter, r *http.Request) {
	var sekolahVar string

	if r.URL.Query().Get("namasekolah") == "" {
		sekolahVar = ""
	} else {
		sekolahVar = r.URL.Query().Get("namasekolah")
	}

	res, err := g.sekolahRepo.GetSearchName(r.Context(), sekolahVar)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}
