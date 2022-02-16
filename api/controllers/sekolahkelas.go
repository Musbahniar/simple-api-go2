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

type SekolahKelas struct {
	sekolahkelasRepo models.SekolahKelasRepo
}

func NewSekolahKelasHandler(db *driver.DB) *SekolahKelas {
	return &SekolahKelas{
		sekolahkelasRepo: r.NewSekolahKelasRepo(db.SQL),
	}
}

func (g *SekolahKelas) GetAll(w http.ResponseWriter, r *http.Request) {
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

	res, err := g.sekolahkelasRepo.GetAll(r.Context(), limit, offset)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *SekolahKelas) GetOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, err := g.sekolahkelasRepo.GetOne(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *SekolahKelas) GetByKelompok(w http.ResponseWriter, r *http.Request) {
	var kelompokVar string

	if r.URL.Query().Get("kelompok") == "" {
		kelompokVar = ""
	} else {
		kelompokVar = r.URL.Query().Get("kelompok")
	}

	res, err := g.sekolahkelasRepo.GetByKelompok(r.Context(), kelompokVar)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}
