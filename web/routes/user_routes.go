package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-l-mathew/GoFileVault/web/handlers"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/", handlers.AuthHandlers).Methods(http.MethodGet)
}
