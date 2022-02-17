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

	"github.com/go-playground/validator/v10"
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
		for _, errVal := range errVal.(validator.ValidationErrors) {
			// messageVal := helper.msgForTag(errVal.Tag())
			arr = append(arr, models.InputError{Field: errVal.Field(), Msg: helper.ValidasiBalik(errVal.Tag())})
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
