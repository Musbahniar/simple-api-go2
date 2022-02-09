package controllers

import (
	"encoding/json"
	"net/http"
	hc "simple-api-go2/api/constants"
	"simple-api-go2/api/models"
	r "simple-api-go2/api/repositories"
	"simple-api-go2/config/driver"
	"simple-api-go2/handler"
)

type Users struct {
	userRepo models.UsersRepo
}

func NewUserHandler(db *driver.DB) *Users {
	return &Users{
		userRepo: r.NewUserRepo(db.SQL),
	}
}

func (u *Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	req := models.User{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	//validate the request
	if req.Name == nil || req.Password == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	res, err := u.userRepo.Create(r.Context(), &req)
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusCreated, res)
}

func (u *Users) VerifyUser(w http.ResponseWriter, r *http.Request) {
	req := models.User{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	//validate the request
	if req.Name == nil || req.Password == nil {
		handler.HttpError(w, http.StatusBadRequest, hc.BAD_REQUEST, err)
		return
	}
	res, err := u.userRepo.Verify(r.Context(), &req)
	if err != nil && err == models.ErrInvalidCredential {
		handler.HttpError(w, http.StatusUnauthorized, err.Error(), err.Error())
		return
	}
	if err != nil {
		handler.HttpError(w, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}
	handler.HttpResponse(w, http.StatusOK, res)
}
