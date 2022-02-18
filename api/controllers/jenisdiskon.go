package controllers

import (
	"encoding/json"
	"net/http"
	hc "simple-api-go2/api/constants"
	"simple-api-go2/api/helper"
	"simple-api-go2/api/models"
	r "simple-api-go2/api/repositories"
	"simple-api-go2/config/driver"
	"simple-api-go2/handler"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type JenisDiskon struct {
	jenisdiksonRepo models.JenisDiskonRepo
}

func NewJenisDiskonHandler(db *driver.DB) *JenisDiskon {
	return &JenisDiskon{
		jenisdiksonRepo: r.NewJenisDiskonRepo(db.SQL),
	}
}

func (g *JenisDiskon) CreateJenisDiskon(w http.ResponseWriter, r *http.Request) {
	jenisdiskon := models.CreateJenisDiskon{}
	err := json.NewDecoder(r.Body).Decode(&jenisdiskon)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	validate := validator.New()
	errVal := validate.Struct(jenisdiskon)
	if errVal != nil {
		var arr []models.InputError
		// Value: fmt.Sprintf("%v", errVal.Value())
		for _, errVal := range errVal.(validator.ValidationErrors) {
			arr = append(arr, models.InputError{Field: errVal.Field(), Msg: helper.ValidasiBalik(errVal.Tag()), Value: errVal.Param()})
		}
		handler.HttpValidasi(w, http.StatusBadRequest, hc.BAD_REQUEST, arr)
		return
	}

	res, err := g.jenisdiksonRepo.Create(r.Context(), &jenisdiskon)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	handler.HttpResponse(w, http.StatusCreated, res)
}

func (g *JenisDiskon) UpdateJenisDiskon(w http.ResponseWriter, r *http.Request) {
	req := models.JenisDiskon{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	validate := validator.New()
	errVal := validate.Struct(req)
	if errVal != nil {
		var arr []models.InputError
		// Value: fmt.Sprintf("%v", errVal.Value())
		for _, errVal := range errVal.(validator.ValidationErrors) {
			arr = append(arr, models.InputError{Field: errVal.Field(), Msg: helper.ValidasiBalik(errVal.Tag()), Value: errVal.Param()})
		}
		handler.HttpValidasi(w, http.StatusBadRequest, hc.BAD_REQUEST, arr)
		return
	}

	res, err := g.jenisdiksonRepo.Update(r.Context(), &req)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *JenisDiskon) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res, err := g.jenisdiksonRepo.Delete(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *JenisDiskon) GetAll(w http.ResponseWriter, r *http.Request) {
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

	res, err := g.jenisdiksonRepo.GetAll(r.Context(), limit, offset)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}

func (g *JenisDiskon) GetOne(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	res, err := g.jenisdiksonRepo.GetOne(r.Context(), vars["id"])
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}
