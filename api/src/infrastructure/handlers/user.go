package handlers

import (
	"api/src/domain/users/service"
	"api/src/infrastructure/mysql"
	inputRequest "api/src/infrastructure/request"
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
		toError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user inputRequest.UserRequest
	if err = json.Unmarshal(request, &user); err != nil {
		toError(w, http.StatusBadRequest, err)
		return
	}

	db, err := mysql.NewMySQLConnection()
	if err != nil {
		toError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	re := mysql.NewUserMySQLRepository(db)
	service := service.NewUserService(re)

	userAdded, err := service.Create(user.CreateDomainUser())

	if err != nil {
		toError(w, http.StatusBadRequest, err)
		return
	}

	toJson(w, http.StatusCreated, userAdded)
}

func NewUsersController() UsersController {
	return UsersController{}
}
