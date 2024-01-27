package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-l-mathew/GoFileVault/web/handlers"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/signup", handlers.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/signin", handlers.SignInUser).Methods(http.MethodPost)
}
