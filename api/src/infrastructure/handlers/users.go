package handlers

import "net/http"

type UsersController struct{}

func (u UsersController) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (u UsersController) Find(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (u UsersController) Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (u UsersController) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (u UsersController) Save(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func NewUsersController() UsersController {
	return UsersController{}
}
