package config

import (
	"log"
	"net/http"
	"simple-api-go2/config/driver"
	"simple-api-go2/handler"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

//Initialize app routing
func (a *App) Intialize() {
	db, err := driver.ConnectDB()
	if err != nil {
		panic(err)
	}
	a.Router = mux.NewRouter()
	handleAppRoutes(a.Router, db)
}

//Run App Here
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(":"+host, handler.CORS(a.Router)))
}
