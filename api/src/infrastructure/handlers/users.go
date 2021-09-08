package handlers

import (
	"api/src/banco"
	"api/src/domain/users"
	"api/src/infrastructure/mysql"
	modelos "api/src/models"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user modelos.Usuario
	if err = json.Unmarshal(request, &user); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := banco.Conectar()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	re := mysql.NewUserMySQLRepository(db)
	service := users.NewUserService(re)

	userAdded, err := service.Create(user.CreateDomainUser())

	if err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, userAdded)
}

func NewUsersController() UsersController {
	return UsersController{}
}
