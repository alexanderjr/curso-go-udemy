package api

import (
	"api/src/infrastructure/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.Headers("content-type", "application/json")

	usersController := handlers.NewUsersController()
	r.HandleFunc("/usuarios", usersController.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/usuarios", usersController.Save).Methods(http.MethodPost)
	r.HandleFunc("/usuarios/{usuarioId}", usersController.Find).Methods(http.MethodGet)
	r.HandleFunc("/usuarios/{usuarioId}", usersController.Update).Methods(http.MethodPatch)
	r.HandleFunc("/usuarios", usersController.Delete).Methods(http.MethodDelete)

	return r
}
