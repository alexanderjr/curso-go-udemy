package controller

import (
	domainService "api/src/domain/users/service"
	"api/src/infrastructure/controller/presenter"
	inputRequest "api/src/infrastructure/controller/request"
	"api/src/infrastructure/mysql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type UsersController struct{}

func (u UsersController) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.Info("Begin UsersController@Save")
	db, err := mysql.NewMySQLConnection()
	if err != nil {
		toError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	re := mysql.NewUserMySQLRepository(db)
	service := domainService.NewUserService(re)

	users, err := service.GetAll()

	if err != nil {
		logger.Error("Err: ", err)
		toError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	toJson(w, http.StatusOK, presenter.ShowAllUser(users))
	logger.Info("End UsersController@Save")
}

func (u UsersController) Find(w http.ResponseWriter, r *http.Request) {
	logger.Info("Begin UsersController@Save")

	parametros := mux.Vars(r)

	userId, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		toError(w, http.StatusBadRequest, erro)
		return
	}

	db, err := mysql.NewMySQLConnection()
	if err != nil {
		toError(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	re := mysql.NewUserMySQLRepository(db)
	service := domainService.NewUserService(re)

	user, err := service.FindById(int(userId))

	if err != nil {
		logger.Error("Err: ", err)

		if errors.Is(err, domainService.ErrUserNotFound) {
			toError(w, http.StatusNotFound, err)
			return
		}

		toError(w, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	toJson(w, http.StatusOK, presenter.ShowUser(*user))
	logger.Info("End UsersController@Save")
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
	service := domainService.NewUserService(re)

	userAdded, err := service.Create(input.CreateDomainUser())

	if err != nil {
		logger.Error("Err: ", err)
		if !service.ErrFromDomain(err) {
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
