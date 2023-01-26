package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Scrackc/api-crud-go/auth"
	"github.com/Scrackc/api-crud-go/model"
)

type login struct {
	storage Storage
}

func newLogin(s Storage) login {
	return login{s}
}

func (l *login) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Login{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "estructura no valida", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	if !isLoginValid(&data) {
		response := newResponse(Error, "credenciales no validas", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	token, err := auth.GenerateToken(&data)
	if err != nil {
		response := newResponse(Error, "No se pudo generar el token", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}
	dataToken := map[string]string{"token": token}
	response := newResponse(Message, "OK", dataToken)
	responseJSON(w, http.StatusOK, response)

}

func isLoginValid(data *model.Login) bool {
	return data.Email == "ed@mail.com" && data.Password == "123456"
}
