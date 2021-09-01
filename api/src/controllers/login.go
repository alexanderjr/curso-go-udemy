package controllers

import (
	"api/src/banco"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"errors"

	// modelos "api/src/models"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	type Login struct {
		Email string `json:"email,omitempty`
		Senha string `json:"senha,omitempty`
	}

	var login Login
	if erro = json.Unmarshal(corpoRequest, &login); erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorEmail(login.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuario.Senha, login.Senha); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, errors.New("invalid credentials"))
		return
	}

	respostas.JSON(w, http.StatusAccepted, login)
}
