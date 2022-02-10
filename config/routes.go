package config

import (
	"encoding/json"
	"net/http"
	"simple-api-go2/api/controllers"
	"simple-api-go2/config/driver"
	"simple-api-go2/middleware"

	"github.com/gorilla/mux"
)

func handleAppRoutes(r *mux.Router, db *driver.DB) {
	//api health check
	r.HandleFunc("/health", healthCheck).Methods(http.MethodGet)

	//handling API versioning
	v1 := r.PathPrefix("/api/v1").Subrouter()

	userHandler := controllers.NewUserHandler(db)
	genreHandler := controllers.NewGenreHandler(db)
	bidangHandler := controllers.NewBidangHandler(db)

	v1.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	v1.HandleFunc("/users/verify", userHandler.VerifyUser).Methods(http.MethodPost)

	v1.HandleFunc("/genre/all", genreHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/genre/{id}", genreHandler.GetOne).Methods(http.MethodGet)

	v1.HandleFunc("/bidang/all", bidangHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/bidang/{id}", bidangHandler.GetOne).Methods(http.MethodGet)

	//Api Key validation middleare for all routes
	v1.Use(middleware.ApiKeyMiddleware)
}

//methos to check api health status
func healthCheck(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response{
		Status: "up",
	})
}
