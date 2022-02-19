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
	tahunajaranHandler := controllers.NewTahunAjaranHandler(db)
	penandaHandler := controllers.NewPenandaHandler(db)
	provinsiHandler := controllers.NewProvinsiHandler(db)
	distrikHandller := controllers.NewDistrikHandler(db)
	sekolahkelasHandler := controllers.NewSekolahKelasHandler(db)
	sekolahHandler := controllers.NewSekolahHandler(db)
	jenisDiskonHandler := controllers.NewJenisDiskonHandler(db)
	daftarDiskonHandler := controllers.NewDaftarDiskonHandler(db)

	v1.HandleFunc("/users", userHandler.CreateUser).Methods(http.MethodPost)
	v1.HandleFunc("/users/verify", userHandler.VerifyUser).Methods(http.MethodPost)

	v1.HandleFunc("/genre/all", genreHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/genre/{id}", genreHandler.GetOne).Methods(http.MethodGet)

	v1.HandleFunc("/bidang/all", bidangHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/bidang/{id}", bidangHandler.GetOne).Methods(http.MethodGet)

	v1.HandleFunc("/tp/all", tahunajaranHandler.GetAll).Methods(http.MethodGet)

	v1.HandleFunc("/penanda/all", penandaHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/penanda/{id}", penandaHandler.GetOne).Methods(http.MethodGet)

	v1.HandleFunc("/provinsi/all", provinsiHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/provinsi/{id}", provinsiHandler.GetOne).Methods(http.MethodGet)

	v1.HandleFunc("/distrik/all", distrikHandller.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/distrik/{id}", distrikHandller.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/distrik/provinsi/{id}", distrikHandller.GetByIdProvinsi).Methods(http.MethodGet)

	v1.HandleFunc("/sekolahkelas/all", sekolahkelasHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/sekolahkelas/{id}", sekolahkelasHandler.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/sekolahkelas/", sekolahkelasHandler.GetByKelompok).Methods(http.MethodGet)

	v1.HandleFunc("/sekolah/all", sekolahHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/sekolah/{id}", sekolahHandler.GetOne).Methods(http.MethodGet)
	v1.HandleFunc("/sekolah/", sekolahHandler.GetSearchName).Methods(http.MethodGet)

	v1.HandleFunc("/jenisdiskon", jenisDiskonHandler.CreateJenisDiskon).Methods(http.MethodPost)
	v1.HandleFunc("/jenisdiskon", jenisDiskonHandler.UpdateJenisDiskon).Methods(http.MethodPut)
	v1.HandleFunc("/jenisdiskon/{id}", jenisDiskonHandler.Delete).Methods(http.MethodDelete)
	v1.HandleFunc("/jenisdiskon/all", jenisDiskonHandler.GetAll).Methods(http.MethodGet)
	v1.HandleFunc("/jenisdiskon/{id}", jenisDiskonHandler.GetOne).Methods(http.MethodGet)

	v1.HandleFunc("/daftardiskon/", daftarDiskonHandler.GetByTahunAjaran).Methods(http.MethodGet)
	v1.HandleFunc("/daftardiskon/{id}", daftarDiskonHandler.GetOne).Methods(http.MethodGet)

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
