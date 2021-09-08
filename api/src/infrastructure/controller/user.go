package controller

import (
	"api/src/domain/users/service"
	"api/src/infrastructure/controller/presenter"
	inputRequest "api/src/infrastructure/controller/request"
	"api/src/infrastructure/mysql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	logger "github.com/sirupsen/logrus"
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
	logger.Info("Begin UsersController@Save")
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		toError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var input inputRequest.UserRequest
	if err = json.Unmarshal(request, &input); err != nil {
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

	userAdded, err := service.Create(input.CreateDomainUser())

	if err != nil {
		if !service.ErrFromDomain(err) {
			logger.Error("Err: ", err)
			toError(w, http.StatusInternalServerError, ErrInternalServer)
			return
		}

		toError(w, http.StatusBadRequest, err)
		return
	}

	toJson(w, http.StatusCreated, presenter.ShowUser(*userAdded))
	logger.Info("End UsersController@Save")
}

func NewUsersController() UsersController {
	return UsersController{}
}
