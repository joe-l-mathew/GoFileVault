package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-l-mathew/GoFileVault/pkg/db"
	"github.com/joe-l-mathew/GoFileVault/web/routes"
)

func main() {
	db.InitDb()
	router := mux.NewRouter()
	routes.UserRoute(router)
	http.ListenAndServe(":8000", router)
}
